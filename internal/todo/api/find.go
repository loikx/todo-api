package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofrs/uuid"
	"todo-api/internal/todo/usecases"
	"todo-api/pkg/errors"
)

type FindHandler struct {
	useCase *usecases.FindUseCase
}

func NewFindHandler(useCase *usecases.FindUseCase) *FindHandler {
	return &FindHandler{useCase: useCase}
}

func (handler *FindHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var findRequest struct {
		Id uuid.UUID `json:"id"`
	}

	err = json.Unmarshal(body, &findRequest)
	if err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusBadRequest)
		return
	}

	response, err := handler.useCase.Handle(request.Context(), findRequest.Id)
	if err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusInternalServerError)
		return
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if _, err = writer.Write(marshalResponse); err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusInternalServerError)
	}
}
