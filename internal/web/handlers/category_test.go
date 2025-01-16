package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/pressly/goose/v3"
)

type Runner struct {
	Connection *db.PostgresConnection
	Handler    *CategoryHandler
}

func NewRunner() *Runner {
	conn, err := db.NewDBConnection(
		"postgres",
		`
host=localhost port=5432 user=postgres 
password=admin dbname=ecommerce_testdb sslmode=disable`)
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

func prepareData() {
	sql := `
delete from public.categories;

insert into public.categories(id, title, code)
values
	('688e64d3-c722-48e5-be96-850e419df2d6', 'category@1', 'cat@1'),
	('996be659-81f0-457c-8682-800abcfd64c2', 'category@2', 'cat@2');`
	if _, err := runner.Connection.DB.Exec(sql); err != nil {
		log.Fatal(err)
	}
}

func start() {
	runner = NewRunner()
	migrate()
	prepareData()
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

func TestGetAll(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	runner.Handler.getAll(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
}

func TestGetByID(t *testing.T) {
	id := "688e64d3-c722-48e5-be96-850e419df2d6"
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf("/categories/%s", id), nil)
	runner.Handler.getByID(rr, req, id)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	out := models.Category{}
	if err := json.NewDecoder(rr.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling category: %v", err)
	}
	if out.Title != "category@1" || out.Code != "cat@1" {
		t.Errorf("Wrong unmarshalling category: %v", out)
	}
}

func TestCreate(t *testing.T) {
	rr := httptest.NewRecorder()
	categoryData := []byte(`{"title":"category@3", "code":"cat@3"}`)
	req := httptest.NewRequest(http.MethodPost, "/categories",
		bytes.NewBuffer(categoryData))
	req.Header.Set("Content-Type", "application/json")
	runner.Handler.create(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusCreated)
	}

	out := models.Category{}
	if err := json.NewDecoder(rr.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling category: %s", err)
	}
	if out.Title != "category@3" || out.Code != "cat@3" {
		t.Errorf("Wrong unmarshalling category, got: %v", out)
	}
}

func TestUpdate(t *testing.T) {
	// ('996be659-81f0-457c-8682-800abcfd64c2', 'category@2', 'cat@2');`
	id := "996be659-81f0-457c-8682-800abcfd64c2"
	rr := httptest.NewRecorder()
	categoryData := []byte(`{"title":"upd title", "code":"upd code"}`)
	req := httptest.NewRequest(http.MethodPut,
		fmt.Sprintf("/categories/%s", id), bytes.NewBuffer(categoryData))
	req.Header.Set("Content-Type", "application/json")
	runner.Handler.update(rr, req, id)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	out := models.Category{}
	if err := json.NewDecoder(rr.Body).Decode(&out); err != nil {
		t.Errorf("Failed to unmarshalling category: %s", err)
	}
	if out.Title != "upd title" || out.Code != "upd code" {
		t.Errorf("Wrong unmarshalling category, got: %v", out)
	}
}
