package main

import (
	"context"
	"fmt"
	"net"

	goal "github.com/anwerj/mygoal/goal"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	fmt.Println("Hi, I am set")

	listener, err := net.Listen("tcp", "localhost:4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	goal.RegisterSelfServiceServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Welcome(ctx context.Context, request *goal.Name) (*goal.WelcomeMessage, error) {
	first, last := request.GetFirst(), request.GetLast()

	result := first + " " + last

	return &goal.WelcomeMessage{Result: result}, nil
}
