package server

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestUpdate(t *testing.T) {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(":1323", opts)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	client := protocol.NewBookServiceClient(conn)
	book := protocol.Book{
		Id: "book1583137384",
		Title: "World and Piece",
		Author: "Lev Tolstoy",
		Genre: "Romance",
		Amount: 2500,
		Year: 1954,
		NumberOfPages: 322,
		Edition: "Moscow-1996",
		IsPopular: false,
		InStock: false,
	}

	_, err = client.Update(context.Background(), &protocol.UpdateRequest{Book: &book})
	if err != nil {
		log.Error(err)
	}

}
