package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xagero/go-cli"
	"github.com/xagero/go-cli/command"
)

func main() {

	// Command with args
	cmd := command.Construct("app:cmd", "Example command")
	cmd.AddArgument("arg1", command.ArgumentRequired, "First argument (*required)")
	cmd.AddArgument("arg2", command.ArgumentOptional, "Second argument (optional)")
	cmd.SetCallback(func() error {

		name := cmd.GetName()
		v1 := cmd.GetArgumentValue("arg1")
		v2 := cmd.GetArgumentValue("arg2")

		fmt.Printf("I am '%s' command with first arg='%s', second arg='%s'\n", name, v1, v2)

		return nil
	})

	// Setup console
	name := "Console"
	description := "Simple console application"
	version := "v0.1"

	console := cli.Construct(name, description, version)
	console.AddCommand(cmd)

	// Run console
	if err := console.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
