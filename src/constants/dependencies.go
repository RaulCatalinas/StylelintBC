package constants

import "stylelintbc/src/types"

var INSTALLATION_COMMANDS = map[types.PackageManager]string{
	"npm":  "install",
	"yarn": "add",
	"pnpm": "add",
	"bun":  "add",
}
