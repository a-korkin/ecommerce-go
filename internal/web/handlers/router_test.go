package handlers

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
)

func TestServeHTTP(t *testing.T) {
	conn, err := db.NewDBConnection("postgres",
		"host=localhost port=5432 user=postgres password=admin dbname=ecommerce_testdb sslmode=disable")
	router := NewRouter(conn.DB)

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
