package api

import (
	"encoding/json"
	"io"
	"net/http"

	"todo-api/internal/todo/usecases"
	"todo-api/pkg/errors"
)

type CreateToDoHandler struct {
	useCase *usecases.CreateUseCase
}

func NewCreateToDoHandler(useCase *usecases.CreateUseCase) *CreateToDoHandler {
	return &CreateToDoHandler{useCase: useCase}
}

func (handler *CreateToDoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusBadRequest)
		return
	}

	command := usecases.CreateCommand{}
	err = json.Unmarshal(body, &command)
	if err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusBadRequest)
		return
	}

	response, err := handler.useCase.Handle(request.Context(), &command)
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
