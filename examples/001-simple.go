package main

import (
	"context"
	"fmt"
	"github.com/xagero/go-cli/command"
)

func main() {

	name := "cmd:simple"
	description := "Simple command"

	cmd := command.Construct(name, description)
	cmd.SetCallback(func() error {
		cmd.PrintHelp()
		return nil
	})

	var args []string
	if err := cmd.Run(context.Background(), args); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
