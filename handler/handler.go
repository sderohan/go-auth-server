package handler

import (
	"net/http"

	"github.com/sderohan/go-auth-server/pkg/config"
	"github.com/sderohan/go-auth-server/service/jwtauth"
)

type IAuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SingOut(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	jwtAuthService *jwtauth.JWTManager
}

func (auth AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
}

func (auth AuthHandler) SingOut(w http.ResponseWriter, r *http.Request) {
}

func (auth AuthHandler) Validate(w http.ResponseWriter, r *http.Request) {
}

func NewAuthHandler() IAuthHandler {
	return AuthHandler{
		jwtAuthService: jwtauth.NewJWTManager(config.GetJWTConfig().SecretKey, config.GetJWTConfig().ExpiryDuration),
	}
}
