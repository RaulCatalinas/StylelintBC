package userinput

import (
	"os"

	"github.com/RaulCatalinas/StylelintBC/internal/enums"
	"github.com/RaulCatalinas/StylelintBC/internal/utils"

	"github.com/AlecAivazis/survey/v2"
)

func GetPackageManager() enums.PackageManager {
	var questions = []*survey.Question{
		{
			Name: "packageManger",
			Prompt: &survey.Select{
				Options: utils.GetPackageManagersAsStrings(),
				Message: "Which package manager do you wanna use?",
				Default: "npm", // <- If we use the enum here the library gives an error.
			},
		},
	}

	answers := struct {
		PackageManager string `survey:"packageManger"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(utils.GetErrorMessage("PackageManagerSelection")),
		})

		os.Exit(1)
	}

	return enums.PackageManager(answers.PackageManager)
}
