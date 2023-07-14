package lib

import (
	"log"

	"github.com/streadway/amqp"
)

type AMQP struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewAMQP(env Env) (*AMQP, error) {
	conn, err := amqp.Dial(env.AMQP)
	if err != nil {
        log.Fatalf("failed dialing amqp: %v\n", err)
	}

	ch, err := conn.Channel()
	if err != nil {
        log.Fatalf("failed opening amqp channel: %v\n", err)
	}

    log.Print("connected to amqp")

	return &AMQP{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *AMQP) Close() error {
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

func (p *AMQP) Publish(msg []byte) error {
    q, err := p.ch.QueueDeclare(
        "video",
        false,
        false,
        false,
        false,
        nil,
        )
	if err != nil {
        return err
	}
    
	err = p.ch.Publish(
        "",
        q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)

    return err
}

func (c *AMQP) Consume() (<-chan amqp.Delivery, error) {
    q, err := c.ch.QueueDeclare(
        "video",
        false,
        false,
        false,
        false,
        nil,
        )
	if err != nil {
        return nil, err
	}

    err = c.ch.Qos(
        1,
        0,
        false,
        )
	if err != nil {
        return nil, err
	}
    
	msgs, err := c.ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
