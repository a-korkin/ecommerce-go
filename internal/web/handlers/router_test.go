package handlers

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
	"github.com/pressly/goose/v3"
)

var router *Router
var connection *db.PostgresConnection

func migrate() {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	migrationDir := filepath.Join(dir, "../../../migrations")
	if err := goose.Up(connection.DB.DB, migrationDir); err != nil {
		log.Fatal(err)
	}
}

func prepareData() {
	sql := `
insert into public.categories(id, title, code)
values
	('688e64d3-c722-48e5-be96-850e419df2d6', 'category@1', 'cat@1'),
	('996be659-81f0-457c-8682-800abcfd64c2', 'category@2', 'cat@2'),
	('efa8b389-a3bd-4e06-84dd-4960a0dfc55b', 'category@3', 'cat@3');`
	if _, err := connection.DB.Exec(sql); err != nil {
		log.Fatal(err)
	}
}

func start() {
	var err error
	connection, err = db.NewDBConnection(
		"postgres",
		`
host=localhost port=5432 user=postgres password=admin
dbname=ecommerce_testdb sslmode=disable`)
	if err != nil {
		log.Fatal(err)
	}
	router = NewRouter(connection.DB)
	migrate()
	prepareData()
}

func shutdown() {
	sql := `delete from public.categories;`
	_, err := connection.DB.DB.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	if err := connection.CloseDBConnection(); err != nil {
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
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	bodyBytes, err := io.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
	}
	want := "hello from main router\n"
	body := string(bodyBytes)
	if body != want {
		t.Errorf("Wrong response body, got: %v, want: %v", body, want)
	}
}
