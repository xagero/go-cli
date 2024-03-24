package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xagero/go-cli"
	"github.com/xagero/go-cli/command"
	"github.com/xagero/go-cli/view"
)

func main() {

	cmd := command.Construct("app:cmd", "Show table view example")
	cmd.DisableCommonOptions()
	cmd.SetCallback(func() error {

		tbl := view.Construct("id", "name", "description", "empty")
		tbl.SetHeading("Simple table view")
		tbl.AddRow(1, "Hello", "Some description at first row")
		tbl.AddRow(2, "World", "Some description at second row")
		tbl.AddRow(3, "!!!", "Some description at third row")
		tbl.Render()

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
