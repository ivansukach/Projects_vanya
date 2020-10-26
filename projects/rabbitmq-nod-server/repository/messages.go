package repository

type Message struct {
	Id    int64 `db:"id"`
	Content string `db:"content"`
}
type Repository interface {
	Create(message *Message) error
	Get(id int64) (*Message, error)
	Update(message *Message) error
	Delete(id int64) error
	Listing() ([]Message, error)
}
