package infrastructure

import (
	"net/http"
	"net/http/httptest"
	domain "task-manager/Domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert" // CORRECTED: Removed the extra "com"
)

// Helper function to create a test JWT token with a consistent secret
func createTestToken(user *domain.User) string {
	// Use a known secret for all tests in this file
	jwtSecret = []byte("test-secret-key")
	token, _ := GenerateJWT(user)
	return token
}

// Helper function to set up a router that mimics the real application's setup for admin routes
func setupRouterForMiddlewareTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// This route group uses both middlewares, just like in router.go
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
	// This tests that AuthMiddleware correctly blocks unauthenticated requests.
	t.Run("No Auth Header on Admin Route", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Authorization header is required"}`, w.Body.String())
	})

	// === Test Case 2: Non-Admin User Accessing Admin Route ===
	// This tests that AdminOnly correctly blocks authenticated non-admins.
	t.Run("Non-Admin User on Admin Route", func(t *testing.T) {
		// Arrange
		user := &domain.User{ID: "user-1", Role: domain.RoleUser}
		token := createTestToken(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.JSONEq(t, `{"error": "Admin access required"}`, w.Body.String())
	})

	// === Test Case 3: Admin User Accessing Admin Route ===
	// This tests the successful case where both middlewares pass.
	t.Run("Admin User on Admin Route", func(t *testing.T) {
		// Arrange
		admin := &domain.User{ID: "admin-1", Role: domain.RoleAdmin}
		token := createTestToken(admin)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "welcome admin"}`, w.Body.String())
	})
}
