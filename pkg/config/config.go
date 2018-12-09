package config

import (
	"github.com/geriBatai/evilsay/pkg/check"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Load(flags *pflag.FlagSet) {
	viper.SetConfigName("evilsay")
	viper.AddConfigPath("$HOME/.config")
	viper.BindPFlags(flags)

	err := viper.ReadInConfig()
	check.WarnOnError(err)
}

func QueueURL() string {
	return viper.Get("amqp_url").(string)
}

func Voice() string {
	return viper.Get("voice").(string)
}
