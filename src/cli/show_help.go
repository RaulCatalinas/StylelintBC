package cli

import (
	"fmt"
	"stylelintbc/src/constants"
)

func ShowHelp() {
	fmt.Println("Usage: stylelintbc [options]")
	fmt.Println()
	fmt.Println("Command line for easy Stylelint configuration")
	fmt.Println()
	fmt.Println("Options:")

	for _, option := range constants.Options {
		fmt.Printf("%-15s %-5s %s\n", option.Option, option.Alias, option.Description)
	}

	fmt.Println()
}
