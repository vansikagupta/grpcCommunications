package main

import (
	"fmt"
	"os"

	"github.com/vansikagupta/grpcCommunications/grpcClient/client"
)

func main() {
	args := os.Args[1:]
	fmt.Println("Starting Client", args[0])
	switch args[0] {
	case "insecure":
		client.InsecureClient()
	case "tls":
		client.TlsClient()
	case "mtls":
		client.MtlsClient()
	default:
		fmt.Println("Incorrect client option")
	}
}
