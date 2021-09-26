package client

import (
	"context"
	"log"

	"github.com/vansikagupta/grpcCommunications/service"
	"google.golang.org/grpc"
)

func InsecureClient() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := service.NewDummyServiceClient(conn)

	response, err := c.SayHello(context.Background(), &service.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
