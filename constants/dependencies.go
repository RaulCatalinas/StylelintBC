package constants

import "stylelintbc/types"

var INSTALLATION_COMMANDS = map[types.PackageManager]string{
	"npm":  "install",
	"yarn": "add",
	"pnpm": "add",
	"bun":  "add",
}
