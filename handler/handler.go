package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sderohan/go-auth-server/pkg/models"
	"github.com/sderohan/go-auth-server/pkg/validator"
)

type IAuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SingOut(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
}

func (auth AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "User registration")
}

func (auth AuthHandler) SingOut(w http.ResponseWriter, r *http.Request) {
	var request models.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("error occured while marshaling the request")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	validationErrors := validator.Validate(request)
	if validationErrors != nil {
		json.NewEncoder(w).Encode(validationErrors)
		return
	}
	json.NewEncoder(w).Encode(request)
}

func (auth AuthHandler) Validate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "token validation")
}

func NewAuthHandler() IAuthHandler {
	return AuthHandler{}
}
