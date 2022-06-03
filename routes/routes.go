package routes

import (
	"github.com/gorilla/mux"
	"github.com/sderohan/go-auth-server/middleware/handler"
)

func InitRoutes(reqHandler handler.IAuthHandler) *mux.Router {
	rts := mux.NewRouter()
	v1 := rts.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/register", reqHandler.RegisterUser).Methods("POST")
	v1.HandleFunc("/generate-token", reqHandler.AuthenticateUser).Methods("POST")
	v1.HandleFunc("/validate-token", reqHandler.ValidateToken).Methods("POST")
	return rts
}
