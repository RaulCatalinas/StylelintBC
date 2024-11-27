package utils

import (
	"fmt"

	"github.com/RaulCatalinas/stylelintbc/internal/types"
)

type WriteMessageProps struct {
	Type    types.MessageType
	Message string
}

func WriteMessage(props WriteMessageProps) {
	switch props.Type {
	case "success":
		fmt.Println("\033[32m" + props.Message + "\033[0m")

	case "info":
		fmt.Println("\033[36m" + props.Message + "\033[0m")

	case "error":
		fmt.Println("\033[31m" + props.Message + "\033[0m")

	case "config":
		fmt.Println("\033[37m" + props.Message + "\033[0m")
	}
}
