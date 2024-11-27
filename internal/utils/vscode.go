package utils

import (
	"encoding/json"
	"os"

	"github.com/RaulCatalinas/stylelintbc/internal/enums"
	"github.com/RaulCatalinas/stylelintbc/internal/types"
)

func ConfigureVSCode() {
	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: "Configuring VSCode...",
	})

	existVSCodeFolder := Exists(".vscode")

	if !existVSCodeFolder {
		CreateFolder(".vscode")
	}

	existVSCodeSettingsFile := Exists(".vscode/settings.json")

	if !existVSCodeSettingsFile {
		CreateEmptyJsonFile(".vscode/settings")
	}

	data, err := os.ReadFile(".vscode/settings.json")

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("VSCodeOrIDXConfig")),
		})

		os.Exit(1)
	}

	var settings types.VSCodeSettings

	deserializeErr := json.Unmarshal(data, &settings)

	if deserializeErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("VSCodeOrIDXConfig")),
		})

		os.Exit(1)
	}

	settings["[css]"] = map[string]interface{}{
		"editor.defaultFormatter": "stylelint.vscode-stylelint",
	}

	jsonData, serializeErr := json.MarshalIndent(settings, "", "  ")

	if serializeErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("VSCodeConfig")),
		})

		os.Exit(1)
	}

	writeFileErr := os.WriteFile(".vscode/settings.json", jsonData, 0644)

	if writeFileErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("VSCodeConfig")),
		})

		os.Exit(1)
	}

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: "VSCode configured successfully",
	})
}

func AddRecommendedExtension(extension string) {
	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: "Adding recommended extensions...",
	})

	existVSCodeFolder := Exists(".vscode")

	if !existVSCodeFolder {
		CreateFolder(".vscode")
	}

	existVSCodeExtensionsFile := Exists(".vscode/extensions.json")

	if !existVSCodeExtensionsFile {
		CreateEmptyJsonFile(".vscode/extensions")
	}

	data, err := os.ReadFile(".vscode/extensions.json")

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("AddRecommendedExtensions")),
		})

		os.Exit(1)
	}

	var extensions types.VSCodeExtensions

	deserializeErr := json.Unmarshal(data, &extensions)

	if deserializeErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("AddRecommendedExtensions")),
		})

		os.Exit(1)
	}

	if extensions.Recommendations == nil {
		extensions.Recommendations = []string{}
	}

	extensions.Recommendations = append(extensions.Recommendations, extension)

	jsonData, serializeErr := json.MarshalIndent(extensions, "", "  ")

	if serializeErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("AddRecommendedExtensions")),
		})

		os.Exit(1)
	}

	writeFileErr := os.WriteFile(".vscode/extensions.json", jsonData, 0644)

	if writeFileErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("AddRecommendedExtensions")),
		})

		os.Exit(1)
	}

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: "Recommended extensions added successfully",
	})
}
