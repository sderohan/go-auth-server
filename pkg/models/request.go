package models

type AuthRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserTokenRequest struct{}

type AuthenticateUserRequest struct{}

type ValidateTokenRequest struct{}
