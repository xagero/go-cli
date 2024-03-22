package cli

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

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

	if len(console.commands) > 0 {
		fmt.Println("\nCommands:")
		for _, cmd := range console.commands {
			fmt.Printf("\t%s \t- %s\n", cmd.GetName(), cmd.GetDescription())
		}
	}
}

// Run execute console command
func (console *Console) Run(context context.Context, args []string) error {

	if len(args) < 2 {
		if console.defaultCommand == nil {
			console.PrintBanner()
			console.PrintHelp()
		} else {
			slog.Debug("Default command call")
			return console.defaultCommand.Run(context, args)
		}
	} else {
		name := args[1]
		if cmd, ok := console.commands[name]; ok {

			a := args[2:]
			console.processArguments(a, cmd)
			console.processOptions(a, cmd)

			cmd.ValidateArgumentRequirement()
			cmd.ValidateOptionRequirement()

			return cmd.Run(context, args)
		} else {

		}
	}

	return nil
}

func (console *Console) processArguments(a []string, cmd *command.Command) {
	idx := -1
	for _, value := range a {
		if strings.HasPrefix(value, "-") {
			continue // skip option
		}

		idx++
		cmd.SetArgumentValue(idx, value)
	}
}

// processOptions
func (console *Console) processOptions(a []string, cmd *command.Command) {
	for _, value := range a {
		if strings.HasPrefix(value, "--") {
			if strings.Contains(value, "=") {
				parts := strings.Split(value, "=")

				k := strings.TrimPrefix(parts[0], "--")
				v := strings.TrimPrefix(parts[1], `"`)

				cmd.SetOptionExists(k, true)
				cmd.SetOptionValue(k, v)

			} else {
				k := strings.TrimPrefix(value, "--")
				cmd.SetOptionExists(k, true)
			}
		} else if strings.HasPrefix(value, "-") {
			short := strings.TrimPrefix(value, "-")
			cmd.SetOptionExistsByShort(short, true)
		}
	}
}
