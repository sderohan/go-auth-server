package routes

import (
	"github.com/gorilla/mux"
	"github.com/sderohan/go-auth-server/middleware/handler"
)

func InitRoutes(reqHandler handler.IAuth) *mux.Router {
	rts := mux.NewRouter()
	v1 := rts.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/authUser", reqHandler.AuthenticateUser)
	v1.HandleFunc("/validateToken", reqHandler.ValidateToken)
	return rts
}
