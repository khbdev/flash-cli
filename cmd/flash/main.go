package main

import (
	"fmt"
	"os"

	"github.com/khbdev/flash-cli/internal/cli"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "flash: %v\n", err)
		os.Exit(1)
	}
}
