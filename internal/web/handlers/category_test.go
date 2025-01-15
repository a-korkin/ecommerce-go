package handlers

import (
	"bytes"

	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
	"github.com/a-korkin/ecommerce/internal/core/services"

	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"path/filepath"

	"github.com/pressly/goose/v3"
)

type Runner struct {
	Connection *db.PostgresConnection
	Handler    *CategoryHandler
}

func NewRunner() *Runner {
	conn, err := db.NewDBConnection(
		"postgres",
		"host=localhost port=5432 user=postgres password=admin dbname=ecommerce_testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	service := services.NewCategoryService(conn.DB)
	return &Runner{
		Connection: conn,
		Handler:    NewCategoryHanlder(service),
	}
}

var runner *Runner

func migrate() {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	migrationDir := filepath.Join(dir, "../../../migrations")
	if err := goose.Up(runner.Connection.DB.DB, migrationDir); err != nil {
		log.Fatal(err)
	}
}

func start() {
	runner = NewRunner()
	migrate()
}

func shutdown(runner *Runner) {
	if err := runner.Connection.CloseDBConnection(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	log.Printf("start main testing...")
	start()
	exitCode := m.Run()
	shutdown(runner)
	log.Printf("stop main testing...")
	os.Exit(exitCode)
}

func TestServeHTTP(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	runner.Handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code, got: %v, want: %v", status, http.StatusOK)
	}

	rr = httptest.NewRecorder()
	newCategory := []byte(`{"title":"test category", "code": "test_cat"}`)
	req = httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(newCategory))
	req.Header.Set("Content-Type", "application/json")
	runner.Handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code, got: %v, want: %v", status, http.StatusCreated)
	}
}
