package cli

import (
	"fmt"
	"stylelintbc/types"
)

func ShowHelp(options []types.Option) {
	fmt.Println("Usage: stylelintbc [options]")
	fmt.Println()
	fmt.Println("Command line for easy Stylelint configuration")
	fmt.Println()
	fmt.Println("Options:")

	for _, option := range options {
		fmt.Printf("%-15s %-5s %s\n", option.Name, option.Alias, option.Description)
	}

	fmt.Println()
}
