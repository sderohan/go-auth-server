package main

import (
	"log"

	"github.com/sderohan/go-auth-server/pkg/config"
	"github.com/sderohan/go-auth-server/routes"
	"github.com/sderohan/go-auth-server/server"
)

func main() {
	config.InitConfigs()
	server := server.NewServer(
		routes.InitRoutes(),
	)
	log.Println("Server listening on", config.GetServerConfig().Address)
	log.Fatal(server.Start())
}
