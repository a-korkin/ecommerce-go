package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from product handler\n"))
}

func main() {
	r := mux.NewRouter()
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	r.HandleFunc("/products", ProductHandler)

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
