package server

import (
	"net/http"

	"github.com/sderohan/go-auth-server/pkg/config"
)

type server struct {
	Handler http.Handler
}

func (srv server) Start() error {

	// get the server config
	srvConfig := config.GetServerConfig()

	server := &http.Server{
		Handler:      srv.Handler,
		Addr:         srvConfig.Address,
		ReadTimeout:  srvConfig.ReadTimeout,
		WriteTimeout: srvConfig.WriteTimeout,
	}

	return server.ListenAndServe()
}

func NewServer(handler http.Handler) server {
	return server{
		Handler: handler,
	}
}
