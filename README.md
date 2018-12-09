# Evilsay

Performs `/usr/bin/say` on another computer.

## Installation

On victim's computer, place `evilsayd` binary into path. Create a configuration file `$HOME/.config/evilsay.yaml` and configure message queue to listen for messages. For example, for locally running RabbitMQ server:

```
amqp_url: http://localhost:15672
```

Start `evilsayd`.


## Running commands

Configure client machine to use the same queue as above. Now you can remotely say any text:

```
$ evilsay "I'm a barbie girl, in a barbie world. Life in plastic is fantastic."
```

`evilsay` also supports custom voices, the same ones that `/usr/bin/say` supports:

```
$ evilsay -v Yuri "Give me your money"
```
