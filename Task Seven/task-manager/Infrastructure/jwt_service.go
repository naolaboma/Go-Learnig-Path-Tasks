package infrastructure

import (
	domain "task-manager/Domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret be loaded from a secure configuration.
var jwtSecret = []byte("your-very-secret-key")

// define the structure of the JWT claims.
type Claims struct {
	UserID   string      `json:"user_id"`
	Username string      `json:"username"`
	Role     domain.Role `json:"role"`
	jwt.RegisteredClaims
}

// create a new JWT token for a given user.
func GenerateJWT(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
