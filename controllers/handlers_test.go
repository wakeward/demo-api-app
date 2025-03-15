package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func MockRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestHealthCheck(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := MockRouter()
	r.GET("/api/v1/healthcheck")

	req, err := http.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)
	if err != nil {
		t.Fatalf("Request failed: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
