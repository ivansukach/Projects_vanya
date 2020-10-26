package service

import "github.com/ivansukach/book-service/repositories"

type BookService struct {
	r repositories.Repository
}

func (bs *BookService) Create(book *repositories.Book) error {
	return bs.r.Create(book)
}
func (bs *BookService) Update(book *repositories.Book) error {
	return bs.r.Update(book)
}
func (bs *BookService) Read(id string) (*repositories.Book, error) {
	return bs.r.Read(id)
}
func (bs *BookService) Delete(id string) error {
	return bs.r.Delete(id)
}
func (bs *BookService) Listing() ([]repositories.Book, error) {
	return bs.r.Listing()
}
func (bs *BookService) DeleteAll() error{
	return bs.r.DeleteAll()
}
func New(repo repositories.Repository) *BookService {
	return &BookService{r: repo}
}
