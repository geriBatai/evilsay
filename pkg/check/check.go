package check

import (
	"fmt"
	"os"
)

func FailOnError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func WarnOnError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Warning: %s\n", err.Error())
	}
}
