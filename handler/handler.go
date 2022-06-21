package handler

import (
	"net/http"
)

type IAuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SingOut(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
}

func (auth AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
}

func (auth AuthHandler) SingOut(w http.ResponseWriter, r *http.Request) {
}

func (auth AuthHandler) Validate(w http.ResponseWriter, r *http.Request) {
}

func NewAuthHandler() IAuthHandler {
	return AuthHandler{}
}
