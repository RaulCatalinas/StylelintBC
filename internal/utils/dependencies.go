package utils

import (
	"os"
	"os/exec"

	"github.com/RaulCatalinas/stylelintbc/internal/constants"
	"github.com/RaulCatalinas/stylelintbc/internal/types"
)

type InstallDependenciesProps struct {
	PackageManagerToUse types.PackageManager
	PackagesToInstall   []string
}

func InstallDependencies(props InstallDependenciesProps) {
	notificationMessage := "Installing dependencies using: " + props.PackageManagerToUse + "..."

	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: string(notificationMessage),
	})

	installationCommand := constants.INSTALLATION_COMMANDS[props.PackageManagerToUse]

	commandArgs := []string{installationCommand}

	commandArgs = append(commandArgs, props.PackagesToInstall...)

	commandArgs = append(commandArgs, "-D")

	cmd := exec.Command(
		string(props.PackageManagerToUse),
		commandArgs...,
	)

	_, err := cmd.CombinedOutput()

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    "error",
			Message: string(GetErrorMessage("Dependencies")),
		})

		os.Exit(1)
	}

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "Dependencies installed successfully",
	})
}
