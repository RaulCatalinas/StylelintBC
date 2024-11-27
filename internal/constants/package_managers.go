package constants

import "github.com/RaulCatalinas/stylelintbc/internal/enums"

var PackageManagers = []enums.PackageManager{
	enums.PackageManagerNpm,
	enums.PackageManagerYarn,
	enums.PackageManagerPnpm,
	enums.PackageManagerBun,
}
