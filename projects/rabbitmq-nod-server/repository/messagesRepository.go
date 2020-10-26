package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type messageRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &messageRepository{db: db}
}

func (mr *messageRepository) Create(message *Message) error {
	_, err := mr.db.NamedExec("INSERT INTO messages (content) VALUES (:content)", message)
	return err
}
func (mr *messageRepository) Get(id int64) (*Message, error) {
	m := Message{}
	err := mr.db.QueryRowx("SELECT * FROM messages WHERE id=$1", id).StructScan(&m)
	return &m, err
}
func (mr *messageRepository) Update(message *Message) error {
	_, err := mr.db.NamedExec("UPDATE messages SET (Content=:content) WHERE id=:id", message)
	return err
}
func (mr *messageRepository) Delete(id int64) error {
	_, err := mr.db.Exec("DELETE FROM messages WHERE id=$1", id)
	return err
}
func (mr *messageRepository) Listing() ([]Message, error) {
	rows, err := mr.db.Queryx("SELECT * FROM messages")
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	m := make([]Message, 0)
	for rows.Next() {
		message := Message{}
		err = rows.StructScan(&message)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		m = append(m, message)
	}
	return m, err
}
