package functions

import "github.com/streadway/amqp"

func RabbitConnect() (*amqp.Connection, *amqp.Channel, amqp.Queue) {

	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")

	HandlerError(err, "RabbitConnect failed")

	ch, err := conn.Channel()

	HandlerError(err, "Channel failed")

	err = ch.ExchangeDeclare(
		"default",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	HandlerError(err, "Declare exchange failed")

	q, err := ch.QueueDeclare(
		"first",
		false,
		false,
		false,
		false,
		nil,
	)

	HandlerError(err, "Queue declare failed")

	err = ch.QueueBind(
		q.Name,
		"first",
		"default",
		false,
		nil,
	)

	HandlerError(err, "QueueBind failed")

	/*defer conn.Close()
	defer ch.Close()*/

	return conn, ch, q
}
