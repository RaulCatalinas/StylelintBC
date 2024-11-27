package userinput

import (
	"os"

	"github.com/RaulCatalinas/stylelintbc/internal/enums"
	"github.com/RaulCatalinas/stylelintbc/internal/utils"

	"github.com/AlecAivazis/survey/v2"
)

func UsingVSCodeEditor() bool {
	var questions = []*survey.Question{
		{
			Name: "usingVSCodeEditor",
			Prompt: &survey.Confirm{
				Message: "Are you using VS Code as a code editor?",
				Default: true,
			},
		},
	}

	answers := struct {
		UsingVSCodeOrIDXEditor bool `survey:"usingVSCodeEditor"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(utils.GetErrorMessage("VSCodeEditor")),
		})

		os.Exit(1)
	}

	return answers.UsingVSCodeOrIDXEditor
}
