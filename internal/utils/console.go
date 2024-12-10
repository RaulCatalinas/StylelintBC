package utils

import (
	"fmt"

	"github.com/RaulCatalinas/StylelintBC/internal/enums"
)

type WriteMessageProps struct {
	Type    enums.MessageType
	Message string
}

func WriteMessage(props WriteMessageProps) {
	switch props.Type {
	case enums.MessageTypeSuccess:
		fmt.Println("\033[32m" + props.Message + "\033[0m")

	case enums.MessageTypeInfo:
		fmt.Println("\033[36m" + props.Message + "\033[0m")

	case enums.MessageTypeError:
		fmt.Println("\033[31m" + props.Message + "\033[0m")

	case enums.MessageTypeConfig:
		fmt.Println("\033[37m" + props.Message + "\033[0m")
	}
}
