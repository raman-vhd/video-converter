package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/raman-vhd/video-converter/internal/lib"
	"github.com/raman-vhd/video-converter/internal/model"
	"github.com/raman-vhd/video-converter/internal/repository"
	"github.com/raman-vhd/video-converter/internal/service"
	"github.com/streadway/amqp"
	"go.uber.org/fx"
)

var Module = fx.Options(
	repository.Module,
	service.Module,
	lib.Module,
)

func main() {
	app := fx.New(
		Module,
		fx.Invoke(bootstrap),
	)
	ctx := context.Background()
	err := app.Start(ctx)
	defer app.Stop(ctx)
	if err != nil {
		log.Fatalf("failed starting app: %v", err)
	}
}

func bootstrap(
	env lib.Env,
	amqpLib *lib.AMQP,
	svc service.IVideoService,
) {
	msgs, err := amqpLib.Consume()
	if err != nil {
		log.Panic(err)
	}

	var forever chan struct{}

	go func() {
		for m := range msgs {
			go func(d amqp.Delivery) {
				var data model.AMQPMsg
				err := json.Unmarshal(d.Body, &data)
				if err != nil {
					log.Println(err)
					return
				}

				err = svc.ConvertVideo(data.VideoID, data.Quality, data.Ext)
				if err != nil {
					log.Println(err)
					return
				}

				d.Ack(false)
			}(m)
		}
	}()

	log.Println("video proccessor running...")

	<-forever
}
