package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
