package userinput

import (
	"os"
	"stylelintbc/src/utils"

	"github.com/AlecAivazis/survey/v2"
)

func ShouldPublishToNPM() bool {
	var questions = []*survey.Question{
		{
			Name: "shouldPublishToNPM",
			Prompt: &survey.Confirm{
				Message: "Will you publish it on npm?",
				Default: false,
			},
		},
	}

	answers := struct {
		ShouldPublishToNPM bool `survey:"shouldPublishToNPM"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    "error",
			Message: string(utils.GetErrorMessage("PublishConfirmation")),
		})

		os.Exit(1)
	}

	return answers.ShouldPublishToNPM
}
