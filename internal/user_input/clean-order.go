package userinput

import (
	"os"

	"github.com/RaulCatalinas/StylelintBC/internal/enums"
	"github.com/RaulCatalinas/StylelintBC/internal/utils"

	"github.com/AlecAivazis/survey/v2"
)

func AddStylelintConfigCleanOrder() bool {
	var questions = []*survey.Question{
		{
			Name: "addCleanOrder",
			Prompt: &survey.Confirm{
				Message: "Do you wanna add stylelint-config-clean-order?",
				Default: true,
			},
		},
	}

	answers := struct {
		AddConfigCleanOrder bool `survey:"addCleanOrder"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(utils.GetErrorMessage("StylelintConfigCleanOrder")),
		})

		os.Exit(1)
	}

	return answers.AddConfigCleanOrder
}
