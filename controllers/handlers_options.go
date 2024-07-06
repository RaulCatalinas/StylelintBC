package controllers

import (
	"os"
	"stylelintbc/constants"
	userInput "stylelintbc/user_input"
	"stylelintbc/utils"
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

	usingVSCodeOrIDXEditor := userInput.UsingVSCodeOrIDXEditor()
	packageManagerToUse := userInput.GetPackageManager()
	addStylelintConfigCleanOrder := userInput.AddStylelintConfigCleanOrder()

	utils.GenerateStylelintConfig(utils.GenerateStylelintConfigProps{
		PackageManagerToUse:          packageManagerToUse,
		UseStylelintConfigCleanOrder: addStylelintConfigCleanOrder,
		UsingVSCodeEditor:            usingVSCodeOrIDXEditor,
	})
}

func HandlerOptionCollaborate() {
	utils.WriteMessage(utils.WriteMessageProps{
		Type:    "info",
		Message: "Opening the GitHub repository...",
	})

	utils.OpenURL(constants.REPOSITORY)
}
