package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {}

func InitRoutes() *mux.Router {
	rts := mux.NewRouter()
	v1 := rts.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/getToken", notImplemented)
	v1.HandleFunc("/validateToken", notImplemented)
	return rts
}
