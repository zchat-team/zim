package main

import (
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/zmicro-team/zim/app/task/internal/app"
	"github.com/zmicro-team/zim/pkg/runtime"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	a := app.New(app.Before(before))

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

func before() error {
	runtime.Setup()
	js := runtime.GetJS()
	if _, err := js.StreamInfo("MSGS"); err != nil {
		if !errors.Is(err, nats.ErrStreamNotFound) {
			return err
		}
		//nats stream add MSGS --subjects "MSGS.*" --ack --max-msgs=-1 --max-bytes=-1 --max-age=-1 --storage file --retention work --max-msg-size=-1 --discard=old
		js.AddStream(&nats.StreamConfig{
			Name:      "MSGS",
			Subjects:  []string{"MSGS.*"},
			Retention: nats.WorkQueuePolicy,
			Storage:   nats.FileStorage,
		})
	}

	if _, err := js.ConsumerInfo("MSGS", "TASK"); err != nil {
		if !errors.Is(err, nats.ErrConsumerNotFound) {
			return err
		}
		//nats consumer add MSGS TASK --filter MSGS.received --ack explicit --pull --deliver all --max-deliver=-1
		_, err := js.AddConsumer("ORDERS", &nats.ConsumerConfig{
			Durable:       "TASK",
			AckPolicy:     nats.AckExplicitPolicy,
			FilterSubject: "MSGS.received",
		})

		if err != nil {
			log.Error(err)
		}
	}
	return nil
}
