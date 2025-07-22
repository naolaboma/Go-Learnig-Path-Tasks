// Infrastructure/auth_middleWare.go
package infrastructure

import (
	"net/http"
	"strings"
	"task-manager/Domain"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

type Claims struct {
	UserID   string      `json:"user_id"`
	Username string      `json:"username"`
	Role     domain.Role `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement auth middleware...
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement admin check...
	}
}
