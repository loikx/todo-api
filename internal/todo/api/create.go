package api

import (
	"encoding/json"
	"io"
	"net/http"

	"todo-api/internal/todo/usecases"
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
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	command := usecases.CreateCommand{}
	err = json.Unmarshal(body, &command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := handler.useCase.Handle(request.Context(), &command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if _, err = writer.Write(marshalResponse); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
