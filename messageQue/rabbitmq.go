package messageQue

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func RabbitConnect() {
	// amqp://username:pw@host
	// 연결
	con, err := amqp.Dial("")
	failOnErr(err, "connect Fail..")

	ch, err := con.Channel()
	failOnErr(err, "fail opne connect Fail..")

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	body := "hi"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plan",
			Body:        []byte(body),
		})

	failOnErr(err, "sened Errro")

	defer con.Close()
}
