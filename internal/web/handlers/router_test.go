package handlers

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/adapters/mock/services"
)

var router *Router

func start() {
	catMock := services.NewCategoryMockService()
	prodMock := services.NewProductsMockService()
	userMock := services.NewUsersMockService()

	router = &Router{
		Categories: NewCategoryHandler(catMock),
		Products:   NewProductHandler(prodMock),
		Users:      NewUserHandler(userMock),
	}
}

func TestMain(m *testing.M) {
	log.Printf("start main testing...")
	start()
	exitCode := m.Run()
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
