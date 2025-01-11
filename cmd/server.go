package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/a-korkin/ecommerce/configs"
	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
	"github.com/a-korkin/ecommerce/internal/web/handlers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	r.HandleFunc("/products", handlers.ProductsHandler)
	r.HandleFunc("/products/{id}", handlers.ProductHandler)

	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conn, err := db.NewDBConnection(
		configs.GetEnv("DB_DRIVER"), configs.GetEnv("DB_CONNECTION"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer func() {
		log.Println("db connection closed")
		if err = conn.CloseDBConnection(); err != nil {
			log.Fatalf("failed to close db connection: %v", err)
		}
	}()

	go func() {
		log.Println("server running")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("server terminated")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err)
	}
	log.Println("shutting down")
}
