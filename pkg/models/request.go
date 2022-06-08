package models

type AuthRequest struct {
	UserName string `json:"username" validate:"required,min=5,max=12"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}

type RegisterUserRequest struct{}

type AuthUserRequest struct {
	UserName string `json:"username" validate:"required,min=4,max=12"`
	Password string `json:"password" validate:"required,min=5,max=12"`
}

type ValidateTokenRequest struct{}
