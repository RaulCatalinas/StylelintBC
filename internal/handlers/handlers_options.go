package handlers

import (
	"os"

	"github.com/RaulCatalinas/stylelintbc/internal/constants"
	userInput "github.com/RaulCatalinas/stylelintbc/internal/user_input"
	"github.com/RaulCatalinas/stylelintbc/internal/utils"
)

func HandlerOptionBuild() {
	existPackageJsonInTheCurrentDirectory := utils.Exists("package.json")

	if !existPackageJsonInTheCurrentDirectory {
		utils.WriteMessage(
			utils.WriteMessageProps{
				Type:    "error",
				Message: string(utils.GetErrorMessage("NotFound")),
			},
		)

		os.Exit(1)
	}

	packageManagerToUse := userInput.GetPackageManager()
	addStylelintConfigCleanOrder := userInput.AddStylelintConfigCleanOrder()
	usingVSCodeOrIDXEditor := userInput.UsingVSCodeOrIDXEditor()

	utils.GenerateStylelintConfig(utils.GenerateStylelintConfigProps{
		PackageManagerToUse:          packageManagerToUse,
		UseStylelintConfigCleanOrder: addStylelintConfigCleanOrder,
		UsingVSCodeOrIDXEditor:       usingVSCodeOrIDXEditor,
	})
}

func HandlerOptionCollaborate() {
	utils.WriteMessage(utils.WriteMessageProps{
		Type:    "info",
		Message: "Opening the GitHub repository...",
	})

	utils.OpenURL(constants.REPOSITORY)
}
