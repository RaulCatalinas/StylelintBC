package cli

import (
	"os"
	"stylelintbc/src/constants"
	"stylelintbc/src/controllers"
)

func ReadUserInput() {
	if len(os.Args) < 2 {
		ShowHelp()

		return
	}

	switch os.Args[1] {
	case constants.Options[3].Option, constants.Options[3].Alias:
		ShowHelp()

	case constants.Options[0].Option, constants.Options[0].Alias:
		ShowVersion()

	case constants.Options[1].Option, constants.Options[1].Alias:
		controllers.HandlerOptionCollaborate()

	case constants.Options[2].Option, constants.Options[2].Alias:
		controllers.HandlerOptionBuild()

	default:
		ShowHelp()
	}
}
