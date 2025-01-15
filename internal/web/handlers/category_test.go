package handlers

import (
	"bytes"

	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
	"github.com/a-korkin/ecommerce/internal/core/services"

	// "github.com/gofrs/uuid"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"path/filepath"

	"github.com/pressly/goose/v3"
)

// func (h *CategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		h.create(w, r)
// 	case "PUT":
// 		path := "/{id}"
// 		vars := utils.GetVars(r.RequestURI, path)
// 		id, ok := vars["id"]
// 		if !ok {
// 			msg := fmt.Sprintf("failed to get id")
// 			http.Error(w, msg, http.StatusBadRequest)
// 			return
// 		}
// 		h.update(w, r, id)
// 	case "GET":
// 		path := "/{id}"
// 		vars := utils.GetVars(r.RequestURI, path)
// 		id, ok := vars["id"]
// 		if ok {
// 			h.getByID(w, r, id)
// 			return
// 		}
// 		h.getAll(w, r)
// 	case "DELETE":
// 		path := "/{id}"
// 		vars := utils.GetVars(r.RequestURI, path)
// 		id, ok := vars["id"]
// 		if ok {
// 			h.delete(w, r, id)
// 		}
// 	}
// }

type Runner struct {
	Connection *db.PostgresConnection
	Handler    *CategoryHandler
}

var runner Runner

func start() {
	runner = Runner{}
	conn, err := db.NewDBConnection(
		"postgres",
		"host=localhost port=5432 user=postgres password=admin dbname=ecommerce_testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	runner.Connection = conn
	service := services.NewCategoryService(runner.Connection.DB)
	runner.Handler = NewCategoryHanlder(service)
	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}
	dir, _ := os.Getwd()
	migrationDir := filepath.Join(dir, "../../../migrations")
	if err = goose.Up(runner.Connection.DB.DB, migrationDir); err != nil {
		log.Fatal(err)
	}
}

func shutdown() {
	if err := runner.Connection.CloseDBConnection(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	log.Printf("start main testing...")
	start()
	exitCode := m.Run()
	shutdown()
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
