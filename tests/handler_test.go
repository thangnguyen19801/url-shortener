package tests

import (
	"github.com/yourusername/url-shortener/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/yourusername/url-shortener/internal/handler"
)

func TestCreateShortURL_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	g := service.NewGenerator(nil)
	re := service.NewRedirect(nil)
	ana := service.NewAnalytics(nil)
	h := handler.NewHandler(g, re, ana)
	r.POST("/api/shorten", h.CreateShortURL)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(`{"url":"not-a-url"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}
