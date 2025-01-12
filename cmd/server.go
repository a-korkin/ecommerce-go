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
	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/web/handlers"
)

func main() {
	conn, err := db.NewDBConnection(
		configs.GetEnv("GOOSE_DRIVER"), configs.GetEnv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer func() {
		log.Println("db connection closed")
		if err = conn.CloseDBConnection(); err != nil {
			log.Fatalf("failed to close db connection: %v", err)
		}
	}()

	server := http.Server{
		Addr: ":8080",
	}

	router := handlers.NewRouter(services.NewProductService(conn.DB))
	http.Handle("/", router)

	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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
