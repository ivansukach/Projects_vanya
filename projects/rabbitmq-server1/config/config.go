package config

import(
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Port             string    `env:"PORT"`
	RabbitMQUrl		 string `env:"RabbitMQUrl"`
}

func Load() (cfg Config) {
	cfg.Port = os.Getenv("PORT")
	cfg.RabbitMQUrl = os.Getenv("RabbitMQUrl")
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return
}
