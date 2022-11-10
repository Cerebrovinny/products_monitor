package main

import (
	"encoding/json"
	"github.com/Cerebrovinny/products_monitor/internal/order/entity"
	amqp "github.com/rabbitmq/amqp091-go"
	"math/rand"
)

func Publish(ch *amqp.Channel, order entity.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func GenerateOrder() entity.Order {
	return entity.Order{
		ID:     uuid.New().String(),
		Prince: rand.Float64(),
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
}
