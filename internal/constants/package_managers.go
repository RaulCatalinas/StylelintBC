package constants

import "github.com/RaulCatalinas/StylelintBC/internal/enums"

var PackageManagers = []enums.PackageManager{
	enums.PackageManagerNpm,
	enums.PackageManagerYarn,
	enums.PackageManagerPnpm,
	enums.PackageManagerBun,
}
