package config

import (
	"fmt"
	"os"

	"stylelintbc/cli"
	"stylelintbc/options"
	"stylelintbc/utils"
)

func ConfigureOptions() {
	options := options.GetOptions()

	if len(os.Args) != 2 {
		cli.ShowHelp(options)

		os.Exit(0)
	}

	for _, option := range options {
		if os.Args[1] == option.Name || os.Args[1] == option.Alias {
			option.Handler()

			os.Exit(0)
		}
	}

	utils.WriteMessage(utils.WriteMessageProps{
		Type:    "error",
		Message: "The option you've tried to execute doesn't exist",
	})

	fmt.Println()

	cli.ShowHelp(options)
}
