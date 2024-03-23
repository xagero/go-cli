package cli

import "fmt"

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
