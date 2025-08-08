package integration

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"hospital-middleware/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPatientSearch(t *testing.T) {
	// Setup test router
	r := gin.Default()
	api.SetupRoutes(r, db) // You'll need to mock your DB

	// Test cases
	t.Run("Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/patient/search", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
