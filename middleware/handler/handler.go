package handler

import (
	"fmt"
	"net/http"
)

type IAuthHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	AuthenticateUser(w http.ResponseWriter, r *http.Request)
	ValidateToken(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
}

func (auth AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "User registration")
}

func (auth AuthHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user authentication")
}

func (auth AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "token validation")
}

func NewAuthHandler() IAuthHandler {
	return AuthHandler{}
}
