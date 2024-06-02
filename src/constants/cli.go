package constants

import (
	"stylelintbc/src/types"
)

var Options = []types.Option{
	{
		Option:      "--version",
		Alias:       "-v",
		Description: "Output the version number",
	},
	{
		Option:      "--collaborate",
		Alias:       "-co",
		Description: "Open GitHub repository for collaboration",
	},
	{
		Option:      "--build",
		Alias:       "-b",
		Description: "Start Stylelint's configuration",
	},
	{
		Option:      "--help",
		Alias:       "-h",
		Description: "Display help for command",
	},
}
