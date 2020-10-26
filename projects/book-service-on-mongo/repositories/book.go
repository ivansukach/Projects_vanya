package repositories

type Book struct {
	Id            string `bson:"Id"`
	Title         string `bson:"Title"`
	Author        string `bson:"Author"`
	Genre         string `bson:"Genre"`
	Edition       string `bson:"Edition"`
	NumberOfPages int32  `bson:"NumberOfPages"`
	Year          int32  `bson:"Year"`
	Amount        int32  `bson:"Amount"`
	IsPopular     bool   `bson:"IsPopular"`
	InStock       bool   `bson:"InStock"`
}
type Repository interface {
	Create(book *Book) error
	Read(id string) (*Book, error)
	Update(book *Book) error
	Delete(id string) error
	DeleteAll() error
	Listing() ([]Book, error)
}
