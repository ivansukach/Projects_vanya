package main

import (
	"context"
	"github.com/ivansukach/kafka-nod-server/repository"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=su password=su dbname=messages sslmode=disable")
	if err != nil {
		log.Error(err)
	}
	rps := repository.New(db)
	defer db.Close()
	topic := "my100topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Error(err)
	}
	defer conn.Close()
	tMes := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-tMes.C:
				log.Println("read")
				conn.SetReadDeadline(time.Now().Add(3 * time.Second))
				batch := conn.ReadBatch(1, 3000) // fetch 10KB min, 1MB max
				for {
					b := make([]byte, 10e3) // 10KB max per message
					amount, err := batch.Read(b)
					if err != nil {
						break
					}
					bb := b[:amount]
					mes := string(bb)
					log.Println(mes)
					currentMessage := repository.Message{Content: mes}
					rps.Create(&currentMessage)
				}
				batch.Close()
			}

	}
}