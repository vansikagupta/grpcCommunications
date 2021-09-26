package service

import (
	"context"
	"fmt"
)

type Service struct {
}

func (s *Service) SayHello(ctx context.Context, clientMsg *Message) (*Message, error) {
	fmt.Println("Message from Client: ", clientMsg.Body)
	return &Message{Body: "Hello from the Server"}, nil
}
