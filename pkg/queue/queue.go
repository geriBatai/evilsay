package queue

import (
	"encoding/json"

	"github.com/geriBatai/evilsay/pkg/check"
	"github.com/streadway/amqp"
)

type Rabbit struct {
	Channel *amqp.Channel
	Queue   *amqp.Queue
	Name    string
}

type Message struct {
	Voice   string `json:"voice"`
	Content string `json:"content"`
}

func New(url string) (*Rabbit, error) {
	retval := &Rabbit{Name: "evilsay"}

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	retval.Channel = ch

	queue, err := ch.QueueDeclare(
		"evilsay",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	retval.Queue = &queue

	return retval, nil

}

func (r *Rabbit) Close() {
	_ = r.Channel.Close()
}

func (r *Rabbit) SendMessage(voice, text string) error {
	msg := Message{Voice: voice, Content: text}
	encoded, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return r.Channel.Publish(
		"",
		r.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        encoded,
		},
	)
}

func (r *Rabbit) Receive(out chan<- Message) {

	msgs, err := r.Channel.Consume(
		r.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	check.WarnOnError(err)

	for d := range msgs {
		var msg = Message{}
		err = json.Unmarshal(d.Body, &msg)
		check.WarnOnError(err)

		out <- msg
	}

}
