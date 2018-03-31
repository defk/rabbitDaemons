package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"runtime"
	"encoding/json"
	"first-rabbit/functions"
	"first-rabbit/handlers"
	"first-rabbit/structures"
)

func main() {

	fmt.Print("Started!\n")

	conn, ch, q := functions.RabbitConnect()

	defer conn.Close()
	defer ch.Close()

	task := make(chan structures.Task)
	out := make(chan string)

	runtime.GOMAXPROCS(3)

	go logger(out)

	go worker(task, out)
	go worker(task, out)
	go worker(task, out)

	reader(ch, q, task)

	fmt.Printf("End!\n")
}

func logger(out <-chan string) {

	for {

		fmt.Printf("%s\n", <-out)
	}
}

func reader(ch *amqp.Channel, q amqp.Queue, task chan<- structures.Task) {

	items, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	functions.HandlerError(err, "Consume channel failed")

	for d := range items {

		var item structures.Task

		err := json.Unmarshal(d.Body, &item)

		functions.HandlerError(err, "Task failed")

		task <- item
	}
}

func worker(task <-chan structures.Task, out chan<- string) {

	var item structures.Task

	for {

		item = <-task

		switch item.Alias {

		case "meteo":
			{

				handlers.Meteo(item, out)
			}
			break
		}
	}
}
