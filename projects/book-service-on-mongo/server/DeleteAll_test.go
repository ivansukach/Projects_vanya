package server

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestDeleteAll(t *testing.T) {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(":1323", opts)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()
	client := protocol.NewBookServiceClient(conn)
	_, err = client.DeleteAll(context.Background(), &protocol.EmptyRequest{})
	if err != nil {
		log.Error(err)
	}
}