package main

import (
	"log"
	"net"

	"github.com/Christian-Placencia/0231673_SistemasDistribuidos/internal/server"
)

func main() {
	// Create server config
	config := &server.Config{}

	// Create gRPC server
	srv, err := server.NewGRPCServer(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	err = srv.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
