package command

import (
	"os/exec"
)

func Say(voice, message string) error {
	cmd := exec.Command("say", "-v", voice, message)
	return cmd.Run()
}
