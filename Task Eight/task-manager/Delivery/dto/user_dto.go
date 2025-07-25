package dto

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PromoteUserRequest struct {
	Username string `json:"username" binding:"required"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
