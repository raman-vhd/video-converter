package lib

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(env Env) (*RabbitMQ, error) {
	conn, err := amqp.Dial(env.AMQP)
	if err != nil {
        log.Fatalf("failed dialing amqp: %v\n", err)
	}

	ch, err := conn.Channel()
	if err != nil {
        log.Fatalf("failed opening amqp channel: %v\n", err)
	}

    log.Print("connected to amqp")

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQ) Close() error {
	if r.ch != nil {
		err := r.ch.Close()
		if err != nil {
			return err
		}
	}

	if r.conn != nil {
		return r.conn.Close()
	}

	return nil
}

