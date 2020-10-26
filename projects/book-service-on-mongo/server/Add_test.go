package server

import (
	"context"
	"fmt"
	"github.com/ivansukach/book-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(":1323", opts)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	client := protocol.NewBookServiceClient(conn)
	book := protocol.Book{
		Id: fmt.Sprintf("book%d", time.Now().Unix()),
		Title: "Captain Blood",
		Author: "Rafael Sabatini",
		Genre: "Realistic Fiction",
		Amount: 2500,
		Year: 1991,
		NumberOfPages: 320,
		Edition: "White Flow",
		IsPopular: false,
		InStock: false,
	}

	_, err = client.Add(context.Background(), &protocol.AddRequest{Book: &book})
	if err != nil {
		log.Error(err)
	}

}
