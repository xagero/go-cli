package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xagero/go-cli"
	"github.com/xagero/go-cli/command"
)

func main() {
	name := "console"
	description := "Simple console application"
	version := "v0.1"

	console := cli.Construct(name, description, version)
	cmd := command.Construct("app:cmd", "First command")
	cmd.AddArgument("arg1", command.ArgumentRequired, "First argument (*required)")
	cmd.SetCallback(func() error {
		fmt.Println("I am first command")
		return nil
	})
	console.AddCommand(cmd)

	console.SetDefaultCommand(cmd)
	if err := console.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
