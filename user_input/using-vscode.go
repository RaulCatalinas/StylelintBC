package userinput

import (
	"os"
	"stylelintbc/utils"

	"github.com/AlecAivazis/survey/v2"
)

func UsingVSCodeEditor() bool {
	var questions = []*survey.Question{
		{
			Name: "usingVSCodeEditor",
			Prompt: &survey.Confirm{
				Message: "Will you use VS Code as a code editor?",
				Default: true,
			},
		},
	}

	answers := struct {
		UsingVSCodeEditor bool `survey:"usingVSCodeEditor"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    "error",
			Message: string(utils.GetErrorMessage("VSCodeEditor")),
		})

		os.Exit(1)
	}

	return answers.UsingVSCodeEditor
}
