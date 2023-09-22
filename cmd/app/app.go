package app

import (
	"log"
	"net/http"
	"savannah/cmd/config"
	"savannah/cmd/db_conn"
	"savannah/internal/handlers"
	"savannah/internal/services"
	"time"

	"github.com/go-chi/chi"
)

type App struct {
	mux           *chi.Mux
	storeServices services.Store
	conf          *config.Conf
}

func StartApp(configPath string) (*App, error) {
	//
	conf, err := config.ReadConfiguration(configPath)
	if err != nil {
		return nil, err
	}

	dbStore, err := db_conn.DatabaseConn(conf)
	if err != nil {
		return nil, err
	}

	store, err := services.NewSQLStore(dbStore, conf)
	if err != nil {
		return nil, err
	}
	allServices := handlers.NewHandlerRepo(store)

	mux := allServices.InitRouter()

	return &App{
		mux:           mux,
		storeServices: store,
		conf:          &conf,
	}, nil
}

func (app *App) ServerConnection() error {
	server := &http.Server{
		Addr:              app.conf.Server.Address,
		Handler:           app.mux,
		ReadTimeout:       time.Duration(app.conf.Server.Timeout) * time.Second,
		ReadHeaderTimeout: time.Duration(app.conf.Server.Timeout) * time.Second,
		WriteTimeout:      time.Duration(app.conf.Server.Timeout) * time.Second,
		IdleTimeout:       time.Duration(app.conf.Server.Timeout) * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %s", err)

	}

	return nil
}
