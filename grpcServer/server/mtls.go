package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"github.com/vansikagupta/grpcCommunications/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func MtlsServer() {
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

	//3. load the client's CA
	pemClientCA, err := ioutil.ReadFile("certs/ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		log.Fatal("failed to add client's CA's certificate")
	}

	// Create the tls config
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // asserts that client must present a valid certificate
		ClientCAs:    certPool,                       // specifies trusted client CAs
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
