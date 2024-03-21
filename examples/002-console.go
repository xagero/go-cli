package main

import (
	"context"
	"fmt"
	"github.com/xagero/go-cli"
	"github.com/xagero/go-cli/command"
	"os"
)

func main() {

	name := "Console"
	description := "Simple console application"
	version := "v0.1"

	console := cli.Construct(name, description, version)

	first := command.Construct("app:first", "First command")
	first.SetCallback(func() error {
		fmt.Println("I am first command")
		return nil
	})
	console.AddCommand(first)

	second := command.Construct("app:second", "Second command")
	second.SetCallback(func() error {
		fmt.Println("I am second command")
		return nil
	})
	console.AddCommand(second)

	console.PrintBanner()
	//console.SetDefaultCommand(first)

	if err := console.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
