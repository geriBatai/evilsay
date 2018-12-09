package main

import (
	"fmt"

	"github.com/geriBatai/evilsay/pkg/check"
	"github.com/geriBatai/evilsay/pkg/command"
	"github.com/geriBatai/evilsay/pkg/config"
	"github.com/geriBatai/evilsay/pkg/queue"
	"github.com/spf13/cobra"
)

var amqpURL string

var rootCmd = &cobra.Command{
	Use:   "evilsayd",
	Short: "Says the message from the queue",
	Run:   runSay,
}

func runSay(cmd *cobra.Command, args []string) {
	var err error

	amqpURL = config.QueueURL()

	mq, err := queue.New(amqpURL)
	check.FailOnError(err)

	messages := make(chan queue.Message)
	fmt.Printf("Starting to listen on the queue...\n")

	go mq.Receive(messages)

	for {
		msg := <-messages
		err = command.Say(msg.Voice, msg.Content)
		check.WarnOnError(err)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&amqpURL, "amqp_url", "u", "", "RabbitMQ URL to read message from")
	config.Load(rootCmd.Flags())
}

func main() {
	err := rootCmd.Execute()
	check.FailOnError(err)
}
