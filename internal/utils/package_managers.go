package utils

import "github.com/RaulCatalinas/stylelintbc/internal/constants"

func GetPackageManagersAsStrings() []string {
	packageManagesAsStrings := make([]string, len(constants.PackageManagers))

	for i, manager := range constants.PackageManagers {
		packageManagesAsStrings[i] = string(manager)
	}

	return packageManagesAsStrings
}
