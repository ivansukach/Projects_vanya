package server

import (
	"context"
	"fmt"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/book-service/repositories"
	"github.com/ivansukach/book-service/service"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	bs *service.BookService
}

func New(bs *service.BookService) *Server {
	return &Server{bs: bs}
}

func (s *Server) Add(ctx context.Context, req *protocol.AddRequest) (*protocol.EmptyResponse, error) {
	book := &repositories.Book{
		Id:            req.Book.Id,
		Title:         req.Book.Title,
		Author:        req.Book.Author,
		Genre:         req.Book.Genre,
		Edition:       req.Book.Edition,
		NumberOfPages: req.Book.NumberOfPages,
		Year:          req.Book.Year,
		Amount:        req.Book.Amount,
		IsPopular:     req.Book.IsPopular,
		InStock:       req.Book.InStock,
	}
	err := s.bs.Create(book)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Delete(ctx context.Context, req *protocol.DeleteRequest) (*protocol.EmptyResponse, error) {
	err := s.bs.Delete(req.Id)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) DeleteAll(ctx context.Context, req *protocol.EmptyRequest) (*protocol.EmptyResponse, error) {
	err := s.bs.DeleteAll()
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Update(ctx context.Context, req *protocol.UpdateRequest) (*protocol.EmptyResponse, error) {
	book := &repositories.Book{
		Id:            req.Book.Id,
		Title:         req.Book.Title,
		Author:        req.Book.Author,
		Genre:         req.Book.Genre,
		Edition:       req.Book.Edition,
		NumberOfPages: req.Book.NumberOfPages,
		Year:          req.Book.Year,
		Amount:        req.Book.Amount,
		IsPopular:     req.Book.IsPopular,
		InStock:       req.Book.InStock,
	}
	err := s.bs.Update(book)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Get(ctx context.Context, req *protocol.GetRequest) (*protocol.GetResponse, error) {
	book, err := s.bs.Read(req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	response := &protocol.Book{
		Id:            book.Id,
		Title:         book.Title,
		Author:        book.Author,
		Genre:         book.Genre,
		Edition:       book.Edition,
		NumberOfPages: book.NumberOfPages,
		Year:          book.Year,
		Amount:        book.Amount,
		IsPopular:     book.IsPopular,
		InStock:       book.InStock,
	}
	return &protocol.GetResponse{Book: response}, nil
}
func (s *Server) Listing(ctx context.Context, req *protocol.EmptyRequest) (*protocol.ListingResponse, error) {
	books, err := s.bs.Listing()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resp := make([]*protocol.Book, 0, 10)
	fmt.Println("ALL BOOKS")
	fmt.Println(books)
	for i := range books {
		fmt.Println("book №", i, " ", books[i].Id)
		temp := new(protocol.Book)
		temp.Id = books[i].Id
		temp.Title = books[i].Title
		temp.Author = books[i].Author
		temp.Genre = books[i].Genre
		temp.Edition = books[i].Edition
		temp.NumberOfPages = books[i].NumberOfPages
		temp.Year = books[i].Year
		temp.Amount = books[i].Amount
		temp.IsPopular = books[i].IsPopular
		temp.InStock = books[i].InStock
		resp = append(resp, temp)
	}
	return &protocol.ListingResponse{Books: resp}, nil
}
