package main

import (
	"log"

	"github.com/sderohan/go-auth-server/handler"
	"github.com/sderohan/go-auth-server/pkg/config"
	"github.com/sderohan/go-auth-server/routes"
	"github.com/sderohan/go-auth-server/server"
)

func main() {
	config.InitConfigs()
	requestHandler := handler.NewAuthHandler()
	server := server.NewServer(
		routes.InitRoutes(requestHandler),
	)
	log.Println("Server listening on", config.GetServerConfig().Address)
	log.Fatal(server.Start())
}
