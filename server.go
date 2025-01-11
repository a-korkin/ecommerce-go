package main

import (
	"context"
	"errors"
	"github.com/a-korkin/ecommerce/internal/web/handlers"
	"log"
	"net/http"
	"os/signal"
	"syscall"

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

	go func() {
		log.Println("server running")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
