package lib

import (
	"os"
)

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	VideoDir   string `mapstructure:"VIDEO_DIR"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	AMQPHost   string `mapstructure:"AMQP_HOST"`
	AMQPPort   string `mapstructure:"AMQP_PORT"`
}

func NewEnv() Env {
	env := Env{
		ServerPort: os.Getenv("SERVER_PORT"),
		VideoDir:   os.Getenv("VIDEO_DIR"),
		DBHost:     os.Getenv("MONGODB_SERVICE_HOST"),
		DBPort:     os.Getenv("MONGODB_SERVICE_PORT"),
		AMQPHost:   os.Getenv("RABBITMQ_SERVICE_HOST"),
		AMQPPort:   os.Getenv("RABBITMQ_SERVICE_PORT"),
	}

	return env
}
