package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xagero/go-cli"
	"github.com/xagero/go-cli/command"
)

func main() {

	// Command with options
	cmd := command.NewCommand("app:cmd", "Example command with options")

	cmd.AddOption("opt1", command.OptionValueNone, "Option without value")
	cmd.AddOption("opt2", command.OptionValueOptional, "Option with optional value")
	cmd.AddOption("opt3", command.OptionValueRequire, "Option with required value")

	cmd.SetCallback(func() error {
		name := cmd.GetName()

		fmt.Printf("\nI am '%s' command, my options:\n", name)
		printListOptions(cmd.ListOptions())

		return nil
	})

	// Setup console
	name := "Console"
	description := "Simple console application"
	version := "v0.1"

	console := cli.Construct(name, description, version)
	console.PrintBanner()
	console.AddCommand(cmd)

	// Run console
	if err := console.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}

// printListOptions print list options
func printListOptions(list map[string]*command.CmdOption) {
	for _, option := range list {

		fmt.Printf("\t'--%s', type:%s, ", option.Name(), option.Input())

		switch option.Input() {
		case command.OptionValueNone:
			if exists := option.Exists(); exists {
				fmt.Println("option exists")
			} else {
				fmt.Println("option NOT exists")
			}
		case command.OptionValueOptional:
			if exists := option.Exists(); exists {
				v := option.Value()
				if v == "" {
					fmt.Printf("option exists, value:'%s'\n", "<none>")
				} else {
					fmt.Printf("option exists, value:'%s'\n", v)
				}
			} else {
				fmt.Println("option NOT exists")
			}
		case command.OptionValueRequire:
			if exists := option.Exists(); exists {
				fmt.Printf("option exists, value:'%s'\n", option.Value())
			} else {
				fmt.Println("option NOT exists")
			}
		}

	}
}
