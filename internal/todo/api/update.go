package api

import (
	"encoding/json"
	"io"
	"net/http"

	"todo-api/internal/todo/usecases"
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
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	command := usecases.UpdateCommand{}
	err = json.Unmarshal(body, &command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.useCase.Handle(request.Context(), &command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
