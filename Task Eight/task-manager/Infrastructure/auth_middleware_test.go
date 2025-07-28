package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"task-manager/Domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createTestToken(user *domain.User) string {
	jwtSecret = []byte("test-secret-key")
	authService := NewAuthService()
	token, _ := authService.GenerateToken(user)
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

	t.Run("No Auth Header on Admin Route", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Non-Admin User on Admin Route", func(t *testing.T) {
		user := &domain.User{ID: "user-1", Role: domain.RoleUser}
		token := createTestToken(user)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("Admin User on Admin Route", func(t *testing.T) {
		admin := &domain.User{ID: "admin-1", Role: domain.RoleAdmin}
		token := createTestToken(admin)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/admin/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
