package cli

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/xagero/go-cli/command"
)

type Console struct {
	name           string
	description    string
	version        string
	defaultCommand *command.Command
	commands       map[string]*command.Command
}

// Construct return new Console
func Construct(name string, description string, version string) *Console {

	console := new(Console)
	console.name = name
	console.description = description
	console.version = version
	console.commands = make(map[string]*command.Command)

	return console
}

// SetDefaultCommand set command by default, return Console
func (console *Console) SetDefaultCommand(command *command.Command) *Console {
	console.defaultCommand = command
	return console
}

// AddCommand add command.Command to Console
func (console *Console) AddCommand(cmd *command.Command) *Console {
	name := cmd.GetName()
	console.commands[name] = cmd
	return console
}

// PrintBanner print banner information
func (console *Console) PrintBanner() {
	fmt.Printf("%s %s - %s\n", console.name, console.version, console.description)
}

// PrintHelp print help command
func (console *Console) PrintHelp() {

	fmt.Println("\nUsage:")
	fmt.Println("\tcommand [arguments] [options]")
	fmt.Println("\nCommands:")
	for _, cmd := range console.commands {
		fmt.Printf("\t%s - %s\n", cmd.GetName(), cmd.GetDescription())
	}
}

// Run execute console command
func (console *Console) Run(context context.Context, args []string) error {

	if len(args) < 2 {
		if console.defaultCommand == nil {
			console.PrintHelp()
		} else {
			slog.Debug("Default command call")
			return console.defaultCommand.Run(context, args)
		}
	} else {
		name := args[1]
		if cmd, ok := console.commands[name]; ok {

			a := args[2:]
			idx := -1
			for _, value := range a {
				idx++
				cmd.SetArgumentValue(idx, value)
			}
			return cmd.Run(context, args)
		} else {

		}
	}

	return nil
}
