package messageQue

import (
	"fmt"
	"log"

	"github.com/kbh0581/techPublish-grpc/configue"
	"github.com/mitchellh/mapstructure"
	"github.com/streadway/amqp"
)

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

type RabbitAuth struct {
	Username string
	Password string
	Hostname string
}

func getAuth() RabbitAuth {
	var authData RabbitAuth
	data := configue.GetConnectionData("rabbitMq")
	err := mapstructure.Decode(data, &authData)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return authData
}

func RabbitConnect() {
	// amqp://username:pw@host
	// 연결
	authData := getAuth()

	con, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s%s/", authData.Username, authData.Password, authData.Hostname))
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
