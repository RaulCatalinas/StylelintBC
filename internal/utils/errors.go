package utils

import "github.com/RaulCatalinas/stylelintbc/internal/constants"

func GetErrorMessage(error string) constants.ErrorMessage {
	return constants.ERROR_MESSAGES[error]
}
