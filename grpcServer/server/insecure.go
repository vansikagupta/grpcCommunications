package server

import (
	"log"
	"net"

	"github.com/vansikagupta/grpcCommunications/service"
	"google.golang.org/grpc"
)

func InsecureServer() {
	//1. set up the address on which to listen for client requests; we are setting up tcp network
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a new server instance with default configurations
	grpcServer := grpc.NewServer()

	// register the service
	server := service.Service{}
	service.RegisterDummyServiceServer(grpcServer, &server)

	// start the server so that it accepts incoming requests on the listener
	grpcServer.Serve(listener)
}
