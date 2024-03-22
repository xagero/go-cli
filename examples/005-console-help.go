package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xagero/go-cli"
	"github.com/xagero/go-cli/command"
)

func main() {

	// Command
	cmd := command.Construct("app:cmd", "Example command, print custom help information")
	cmd.DisableCommonOptions()

	help := cmd.AddOption("help", command.OptionValueNone, "Show command help")
	help.SetShortSyntax("h")

	cmd.SetCallback(func() error {
		cmd.PrintHelp() // <-- your help here
		return nil
	})

	name := "Console"
	desc := "Simple console application"
	version := "v0.1"

	console := cli.Construct(name, desc, version)
	console.AddCommand(cmd)

	if err := console.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
