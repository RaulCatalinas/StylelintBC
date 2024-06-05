package utils

import (
	"encoding/json"
	"os"
	"stylelintbc/types"
)

type GenerateStylelintConfigProps struct {
	PackageManagerToUse          types.PackageManager
	UsingVSCodeEditor            bool
	UseStylelintConfigCleanOrder bool
}

func GenerateStylelintConfig(props GenerateStylelintConfigProps) {
	WriteMessage(WriteMessageProps{
		Type:    "config",
		Message: "Generating Stylelint's configuration...",
	})

	packagesToInstall := []string{"stylelint", "stylelint-config-standard"}

	if props.UseStylelintConfigCleanOrder {
		packagesToInstall = append(packagesToInstall, "stylelint-config-clean-order")
	}

	InstallDependencies(InstallDependenciesProps{
		PackageManagerToUse: props.PackageManagerToUse,
		PackagesToInstall:   packagesToInstall,
	})

	createStylelintConfigFiles(props.UseStylelintConfigCleanOrder)

	if props.UsingVSCodeEditor {
		ConfigureVSCode()
		AddRecommendedExtension("stylelint.vscode-stylelint")
	}

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "Stylelint's configuration generated successfully",
	})
}

func createStylelintConfigFiles(useStylelintConfigCleanOrder bool) {
	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: "Creating configuration file...",
	})

	existsStylelintrcFile := Exists(".stylelintrc.json")

	if !existsStylelintrcFile {
		CreateEmptyJsonFile(".stylelintrc")
	}

	data, err := os.ReadFile(".stylelintrc.json")

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    "error",
			Message: string(GetErrorMessage("ConfigFilesCreate")),
		})

		os.Exit(1)
	}

	var stylelintConfig types.StylelintConfig

	deserializeErr := json.Unmarshal(data, &stylelintConfig)

	if deserializeErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    "error",
			Message: string(GetErrorMessage("ConfigFilesCreate")),
		})

		os.Exit(1)
	}

	if stylelintConfig.Extends == nil {
		stylelintConfig.Extends = []string{}
	}

	stylelintConfigPackages := []string{"stylelint-config-standard"}

	if useStylelintConfigCleanOrder {
		stylelintConfigPackages = append(stylelintConfigPackages, "stylelint-config-clean-order")
	}

	stylelintConfig.Extends = append(stylelintConfig.Extends, stylelintConfigPackages...)

	jsonData, serializeErr := json.MarshalIndent(stylelintConfig, "", "  ")

	if serializeErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    "error",
			Message: string(GetErrorMessage("ConfigFilesCreate")),
		})

		os.Exit(1)
	}

	writeFileErr := os.WriteFile(".stylelintrc.json", jsonData, 0644)

	if writeFileErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    "error",
			Message: string(GetErrorMessage("ConfigFilesCreate")),
		})

		os.Exit(1)
	}

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "Configuration file (.stylelintrc.json) created successfully",
	})
}
