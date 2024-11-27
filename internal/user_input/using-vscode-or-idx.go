package userinput

import (
	"os"

	"github.com/RaulCatalinas/stylelintbc/internal/utils"

	"github.com/AlecAivazis/survey/v2"
)

func UsingVSCodeOrIDXEditor() bool {
	var questions = []*survey.Question{
		{
			Name: "usingVSCodeOrIDXEditor",
			Prompt: &survey.Confirm{
				Message: "Are you using VS Code or IDX as a code editor?",
				Default: true,
			},
		},
	}

	answers := struct {
		UsingVSCodeOrIDXEditor bool `survey:"usingVSCodeOrIDXEditor"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    "error",
			Message: string(utils.GetErrorMessage("VSCodeEditor")),
		})

		os.Exit(1)
	}

	return answers.UsingVSCodeOrIDXEditor
}
