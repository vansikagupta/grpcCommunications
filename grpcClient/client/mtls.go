package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/vansikagupta/grpcCommunications/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func MtlsClient() {

	// 1. load trusted server CA
	pemServerCA, err := ioutil.ReadFile("certs/ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		log.Fatal("failed to add server CA's certificate")
	}

	// 2. load client cert and key
	clientCert, certErr := tls.LoadX509KeyPair("certs/client-cert.pem", "certs/client-key.pem")
	if certErr != nil {
		log.Fatalf("failed to load certs: %v", certErr)
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredentials := credentials.NewTLS(config)

	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	defer conn.Close()

	c := service.NewDummyServiceClient(conn)

	response, err := c.SayHello(context.Background(), &service.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
