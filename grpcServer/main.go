package main

import (
	"fmt"
	"os"

	"github.com/vansikagupta/grpcCommunications/grpcServer/server"
)

func main() {
	args := os.Args[1:]
	fmt.Println("Starting Server", args[0])
	switch args[0] {
	case "insecure":
		server.InsecureServer()
	case "tls":
		server.TlsServer()
	case "mtls":
		server.MtlsServer()
	default:
		fmt.Println("Incorrect server option")
	}
}
