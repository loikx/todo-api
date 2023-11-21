package todo

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"todo-api/internal/todo/api"
	"todo-api/internal/todo/config"
	"todo-api/internal/todo/domain"
	"todo-api/internal/todo/repository"
	"todo-api/internal/todo/usecases"
	"todo-api/pkg/server"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type Server interface {
	Serve() error
}

type App struct {
	config *config.Config

	server Server

	router http.Handler

	createHandler *api.CreateToDoHandler
	updateHandler *api.UpdateToDoHandler
	findHandler   *api.FindHandler

	createUseCase *usecases.CreateUseCase
	updateUseCase *usecases.UpdateUseCase
	findUseCase   *usecases.FindUseCase

	repository domain.ToDoRepository

	connection *pgx.Conn
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() error {
	log.Printf("server starts on: %s:%d\n", a.config.Address, a.config.Port)
	return a.server.Serve()
}

func (a *App) Init(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("config load error: %w", err)
	}

	a.config = cfg

	if err = a.initServer(ctx); err != nil {
		return fmt.Errorf("init server: %w", err)
	}

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	con, err := pgx.Connect(ctx, a.config.DataBaseURL)
	if err != nil {
		return fmt.Errorf("init database: %w", err)
	}

	a.connection = con

	a.repository = repository.NewRepository(a.connection)

	a.createUseCase = usecases.NewCreateUseCase(a.repository)
	a.updateUseCase = usecases.NewUpdateUseCase(a.repository)
	a.findUseCase = usecases.NewFindUseCase(a.repository)

	a.createHandler = api.NewCreateToDoHandler(a.createUseCase)
	a.updateHandler = api.NewUpdateToDoHandler(a.updateUseCase)
	a.findHandler = api.NewFindHandler(a.findUseCase)

	a.server = server.NewServer(fmt.Sprintf(
		"%s:%d", a.config.Address, a.config.Port),
		a.createRouter(),
	)

	return nil
}

func (a *App) createRouter() http.Handler {
	router := mux.NewRouter()

	router.Handle("/api/todo/find", a.findHandler).Methods(http.MethodPost)
	router.Handle("/api/todo/create", a.createHandler).Methods(http.MethodPost)
	router.Handle("/api/todo/update", a.updateHandler).Methods(http.MethodPost)

	return router
}
