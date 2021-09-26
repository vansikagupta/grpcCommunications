package server

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/vansikagupta/grpcCommunications/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TlsServer() {
	//1. set up the address on which to listen for client requests; we are setting up tcp network
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//2. Load server's certificate and private key
	serverCert, certErr := tls.LoadX509KeyPair("certs/server-cert.pem", "certs/server-key.pem")
	if certErr != nil {
		log.Fatalf("failed to load certs: %v", certErr)
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert, // asserts that client need not present certificate
	}
	tlsCreds := credentials.NewTLS(config)

	// create a new server instance
	grpcServer := grpc.NewServer(grpc.Creds(tlsCreds))

	// register the service
	server := service.Service{}
	service.RegisterDummyServiceServer(grpcServer, &server)

	// start the server so that it accepts incoming requests on the listener
	grpcServer.Serve(listener)
}
