package main

import (
	"fmt"
	"github.com/xagero/go-cli"
)

func main() {

	name := "cmd"
	description := "Simple cmd command"

	cmd := cli.Construct(name, description)
	cmd.SetCallback(func() error {
		cmd.PrintHelp()
		return nil
	})

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
