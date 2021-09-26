
serverInsecure:
	@go run grpcServer/main.go insecure

serverTls:
	@go run grpcServer/main.go tls

serverMtls:
	@go run grpcServer/main.go mtls

clientInsecure:
	@go run grpcClient/main.go insecure

clientTls:
	@go run grpcClient/main.go tls

clientMtls:
	@go run grpcClient/main.go mtls