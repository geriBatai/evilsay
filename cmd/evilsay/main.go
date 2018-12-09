package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/geriBatai/evilsay/pkg/check"
	"github.com/geriBatai/evilsay/pkg/config"
	"github.com/geriBatai/evilsay/pkg/queue"
)

var voice string
var message string
var amqpURL string

var rootCmd = &cobra.Command{
	Use:   "evilsay",
	Short: "Say something evil",
	Run:   runEvilSay,
}

func runEvilSay(cmd *cobra.Command, args []string) {
	message = strings.Join(args, " ")
	fmt.Printf("[%s]: %s\n", voice, message)

	amqpURL = config.QueueURL()
	voice = config.Voice()

	mq, err := queue.New(amqpURL)
	check.FailOnError(err)

	err = mq.SendMessage(voice, message)
	check.FailOnError(err)

	mq.Close()
}

func init() {
	rootCmd.Flags().StringVarP(&voice, "voice", "v", "Alex", "voice to say a message with")
	rootCmd.Flags().StringVarP(&amqpURL, "amqp_url", "u", "", "RabbitMQ to send command to")

	config.Load(rootCmd.Flags())
}

func main() {
	err := rootCmd.Execute()
	check.FailOnError(err)
}
