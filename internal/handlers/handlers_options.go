package handlers

import (
	"os"

	"github.com/RaulCatalinas/StylelintBC/internal/constants"
	"github.com/RaulCatalinas/StylelintBC/internal/enums"
	userInput "github.com/RaulCatalinas/StylelintBC/internal/user_input"
	"github.com/RaulCatalinas/StylelintBC/internal/utils"
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
