package lib

import (
	"log"

	"github.com/spf13/viper"
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
	env := Env{}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read configuration: %v\n", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("️failed to load environment variables: %v\n", err)
	}

	return env
}
