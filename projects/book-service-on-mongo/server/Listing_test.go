package server

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestListing(t *testing.T) {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(":1323", opts)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()
	client := protocol.NewBookServiceClient(conn)
	response, err := client.Listing(context.Background(), &protocol.EmptyRequest{})
	if err != nil {
		log.Error(err)
	}
	for i := range response.Books {
		log.Println(response.Books[i])
	}
}
