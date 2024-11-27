package utils

import (
	"os"
	"os/exec"

	"github.com/RaulCatalinas/stylelintbc/internal/constants"
	"github.com/RaulCatalinas/stylelintbc/internal/enums"
)

type InstallDependenciesProps struct {
	PackageManagerToUse enums.PackageManager
	PackagesToInstall   []string
}

func InstallDependencies(props InstallDependenciesProps) {
	notificationMessage := "Installing dependencies using: " + props.PackageManagerToUse + "..."

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
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
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("Dependencies")),
		})

		os.Exit(1)
	}

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: "Dependencies installed successfully",
	})
}
