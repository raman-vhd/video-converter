package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DB         string `mapstructure:"DB"`
	AMQP       string `mapstructure:"AMQP"`
	VideoDir   string `mapstructure:"VIDEO_DIR"`
}

func NewEnv() Env {
	var env Env

	viper.SetConfigFile("config/config.env")
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
