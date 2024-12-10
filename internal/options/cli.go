package options

import (
	"github.com/RaulCatalinas/StylelintBC/internal/cli"
	"github.com/RaulCatalinas/StylelintBC/internal/handlers"
	"github.com/RaulCatalinas/StylelintBC/internal/types"
)

func GetOptions() []types.Option {
	return []types.Option{
		{
			Name:        "--version",
			Alias:       "-v",
			Description: "Output the version number",
			Handler:     cli.ShowVersion,
		},
		{
			Name:        "--collaborate",
			Alias:       "-co",
			Description: "Open GitHub repository for collaboration",
			Handler:     handlers.HandlerOptionCollaborate,
		},
		{
			Name:        "--build",
			Alias:       "-b",
			Description: "Start Stylelint's configuration",
			Handler:     handlers.HandlerOptionBuild,
		},
		{
			Name:        "--help",
			Alias:       "-h",
			Description: "Display help for command",
			Handler: func() {
				cli.ShowHelp(GetOptions())
			},
		},
	}
}
