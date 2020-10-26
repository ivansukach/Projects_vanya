package server

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestGet(t *testing.T) {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(":1323", opts)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()
	client := protocol.NewBookServiceClient(conn)
	id := "book1583137481"
	response, err := client.Get(context.Background(), &protocol.GetRequest{Id: id})
	if err != nil {
		log.Error(err)
	}
	log.Println(response.Book)
}
