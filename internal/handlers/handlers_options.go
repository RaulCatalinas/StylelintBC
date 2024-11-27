package handlers

import (
	"os"

	"github.com/RaulCatalinas/stylelintbc/internal/constants"
	"github.com/RaulCatalinas/stylelintbc/internal/enums"
	userInput "github.com/RaulCatalinas/stylelintbc/internal/user_input"
	"github.com/RaulCatalinas/stylelintbc/internal/utils"
)

func HandlerOptionBuild() {
	existPackageJsonInTheCurrentDirectory := utils.Exists("package.json")

	if !existPackageJsonInTheCurrentDirectory {
		utils.WriteMessage(
			utils.WriteMessageProps{
				Type:    enums.MessageTypeError,
				Message: string(utils.GetErrorMessage("NotFound")),
			},
		)

		os.Exit(1)
	}

	packageManagerToUse := userInput.GetPackageManager()
	addStylelintConfigCleanOrder := userInput.AddStylelintConfigCleanOrder()
	usingVSCodeEditor := userInput.UsingVSCodeEditor()

	utils.GenerateStylelintConfig(utils.GenerateStylelintConfigProps{
		PackageManagerToUse:          packageManagerToUse,
		UseStylelintConfigCleanOrder: addStylelintConfigCleanOrder,
		UsingVSCodeEditor:            usingVSCodeEditor,
	})
}

func HandlerOptionCollaborate() {
	utils.WriteMessage(utils.WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: "Opening the GitHub repository...",
	})

	utils.OpenURL(constants.REPOSITORY)
}
