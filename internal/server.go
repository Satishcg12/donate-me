package internal

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/satishcg12/donate-me/internal/db"
)

type App struct {
	router http.Handler
	db     *sql.DB
}

func NewApp() *App {
	db, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database")

	app := &App{
		db: db,
	}
	app.LoadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	defer func() {
		if err := a.db.Close(); err != nil {
			fmt.Println("Error closing the database connection")
		}
		fmt.Println("Database connection closed")
	}()

	fmt.Println("Server is running on port 3000")
	ch := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("server error: %w", err)
		}
		close(ch)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Shutting down server")
		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return server.Shutdown(timeout)
	case err := <-ch:
		return err
	}

}
