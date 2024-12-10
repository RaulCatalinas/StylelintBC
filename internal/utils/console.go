package utils

import (
	"fmt"

	"github.com/RaulCatalinas/StylelintBC/internal/constants"
	"github.com/RaulCatalinas/StylelintBC/internal/enums"
)

type WriteMessageProps struct {
	Type    enums.MessageType
	Message string
}

func WriteMessage(props WriteMessageProps) {
	color, exists := constants.MESSAGE_COLORS[props.Type]

	if !exists {
		color = constants.DEFAULT_COLOR
	}

	fmt.Println(color + props.Message + constants.DEFAULT_COLOR)
}
