package repositories

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(collection *mongo.Collection) Repository {
	return &bookRepository{ collection:collection}
}

type bookRepository struct {
	collection *mongo.Collection
}
func NewMongoClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (br *bookRepository) Create(book *Book) error {
	_, err := br.collection.InsertOne(context.TODO(), book)
	return err
}
func (br *bookRepository) Read(id string) (*Book, error) {
	u := Book{}
	filter := bson.D{{"Id", id}}
	err := br.collection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &u, err
}
func (br *bookRepository) Update(book *Book) error {
	filter := bson.D{{"Id", book.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"Amount", book.Amount},
			{"Author", book.Author},
			{"Edition", book.Edition},
			{"Genre", book.Genre},
			{"InStock", book.InStock},
			{"IsPopular", book.IsPopular},
			{"NumberOfPages", book.NumberOfPages},
			{"Title", book.Title},
			{"Year", book.Year},
		}},
	}
	_, err := br.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func (br *bookRepository) Delete(id string) error {
	filter := bson.D{{"Id", id}}
	_, err := br.collection.DeleteOne(context.TODO(), filter)
	return err
}
func (br *bookRepository) DeleteAll() error {
	filter := bson.D{}
	_, err := br.collection.DeleteMany(context.TODO(), filter)
	return err
}
func (br *bookRepository) Listing() ([]Book, error) {
	books := make([]Book, 0)
	filter := bson.M{}
	cursor, err := br.collection.Find(context.TODO(), filter) //SELECT *?
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		book := new(Book)
		err = cursor.Decode(&book)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		books = append(books, *book)
	}
	return books, err
}
