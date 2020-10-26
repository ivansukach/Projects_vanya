package main

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/book-service/repositories"
	"github.com/ivansukach/book-service/server"
	"github.com/ivansukach/book-service/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	client:=repositories.NewMongoClient()
	err := client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("books").Collection("books")

	rps := repositories.New(collection)
	bs := service.New(rps)
	srv := server.New(bs)
	s := grpc.NewServer()
	protocol.RegisterBookServiceServer(s, srv)
	listener, err := net.Listen("tcp", ":1323")
	if err != nil {
		log.Error(err)
		return
	}
	s.Serve(listener)
}
