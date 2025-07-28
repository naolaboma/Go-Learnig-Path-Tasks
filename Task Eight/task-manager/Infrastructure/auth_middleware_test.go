package infrastructure

import (
	"net/http"
	"net/http/httptest"
	domain "task-manager/Domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createTestToken(user *domain.User) string {
	jwtSecret = []byte("test-secret-key")
	token, _ := GenerateJWT(user)
	return token
}

func setupRouterForMiddlewareTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(AuthMiddleware())
	adminRoutes.Use(AdminOnly())
	{
		adminRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "welcome admin"})
		})
	}
	return r
}

func TestMiddlewareIntegration(t *testing.T) {
	router := setupRouterForMiddlewareTest()

	// === Test Case 1: No Authorization Header ===
	t.Run("No Auth Header on Admin Route", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Authorization header is required"}`, w.Body.String())
	})

	// === Test Case 2: Non-Admin User Accessing Admin Route ===
	t.Run("Non-Admin User on Admin Route", func(t *testing.T) {
		user := &domain.User{ID: "user-1", Role: domain.RoleUser}
		token := createTestToken(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.JSONEq(t, `{"error": "Admin access required"}`, w.Body.String())
	})

	// === Test Case 3: Admin User Accessing Admin Route ===
	t.Run("Admin User on Admin Route", func(t *testing.T) {
		// Arrange
		admin := &domain.User{ID: "admin-1", Role: domain.RoleAdmin}
		token := createTestToken(admin)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "welcome admin"}`, w.Body.String())
	})
}
