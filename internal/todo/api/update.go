package api

import (
	"encoding/json"
	"io"
	"net/http"

	"todo-api/internal/todo/usecases"
	"todo-api/pkg/errors"
)

type UpdateToDoHandler struct {
	useCase *usecases.UpdateUseCase
}

func NewUpdateToDoHandler(useCase *usecases.UpdateUseCase) *UpdateToDoHandler {
	return &UpdateToDoHandler{useCase: useCase}
}

func (handler *UpdateToDoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusBadRequest)
		return
	}

	command := usecases.UpdateCommand{}
	err = json.Unmarshal(body, &command)
	if err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusBadRequest)
		return
	}

	err = handler.useCase.Handle(request.Context(), &command)
	if err != nil {
		bytes, _ := json.Marshal(errors.Error{Message: err.Error()})
		http.Error(writer, string(bytes), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
