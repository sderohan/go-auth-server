package routes

import (
	"github.com/gorilla/mux"
	"github.com/sderohan/go-auth-server/handler"
)

func InitRoutes(reqHandler handler.IAuthHandler) *mux.Router {
	rts := mux.NewRouter()
	v1 := rts.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/sign-in", reqHandler.SignIn).Methods("POST")
	v1.HandleFunc("/sign-out", reqHandler.SingOut).Methods("POST")
	v1.HandleFunc("/validate", reqHandler.Validate).Methods("POST")
	return rts
}
