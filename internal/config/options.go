package config

import (
	"fmt"
	"os"

	"github.com/RaulCatalinas/stylelintbc/internal/cli"
	"github.com/RaulCatalinas/stylelintbc/internal/enums"
	"github.com/RaulCatalinas/stylelintbc/internal/options"
	"github.com/RaulCatalinas/stylelintbc/internal/utils"
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
		Type:    enums.MessageTypeError,
		Message: "The option you've tried to execute doesn't exist",
	})

	fmt.Println()

	cli.ShowHelp(options)
}
