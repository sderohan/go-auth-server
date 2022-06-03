package handler

import "net/http"

type IAuth interface {
	AuthenticateUser(w http.ResponseWriter, r *http.Request)
	ValidateToken(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
}

func (auth AuthHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {

}

func (auth AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
}

func NewAuthHandler() IAuth {
	return AuthHandler{}
}
