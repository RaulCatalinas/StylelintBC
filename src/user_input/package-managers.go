package userinput

import (
	"os"
	"stylelintbc/src/constants"
	"stylelintbc/src/types"
	"stylelintbc/src/utils"

	"github.com/AlecAivazis/survey/v2"
)

func GetPackageManager() types.PackageManager {
	var questions = []*survey.Question{
		{
			Name: "packageManger",
			Prompt: &survey.Select{
				Options: constants.PackageManagers,
				Message: "Which package manager do you wanna use?",
				Default: "npm",
			},
		},
	}

	answers := struct {
		PackageManager string `survey:"packageManger"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    "error",
			Message: string(utils.GetErrorMessage("PackageManagerSelection")),
		})

		os.Exit(1)
	}

	return types.PackageManager(answers.PackageManager)
}
