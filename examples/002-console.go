package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xagero/go-cli"
	"github.com/xagero/go-cli/command"
)

func main() {

	first := command.Construct("app:first", "First command")
	first.DisableCommonOptions()
	first.SetCallback(func() error {
		fmt.Println("I am first command")
		return nil
	})

	second := command.Construct("app:second", "Second command")
	second.DisableCommonOptions()
	second.SetCallback(func() error {
		fmt.Println("I am second command")
		return nil
	})

	name := "Console"
	desc := "Simple console application"
	version := "v0.1"

	console := cli.Construct(name, desc, version)
	console.AddCommand(first)
	console.AddCommand(second)

	if err := console.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
