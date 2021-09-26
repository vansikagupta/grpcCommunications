package main

import (
	"fmt"

	"github.com/vansikagupta/grpcCommunications/server"
)

func main() {
	fmt.Println("Starting Server")
	//server.InsecureServer()
	server.TlsServer()
}
