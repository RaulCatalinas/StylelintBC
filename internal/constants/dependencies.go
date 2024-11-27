package constants

import "github.com/RaulCatalinas/stylelintbc/internal/types"

var INSTALLATION_COMMANDS = map[types.PackageManager]string{
	"npm":  "install",
	"yarn": "add",
	"pnpm": "add",
	"bun":  "add",
}
