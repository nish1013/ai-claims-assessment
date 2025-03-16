package api

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
    router := gin.Default()
    RegisterRoutes(router)

    req, _ := http.NewRequest("GET", "/api/v1/health", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.JSONEq(t, `{"status":"ok"}`, w.Body.String())
}
