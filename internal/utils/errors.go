package utils

import "github.com/RaulCatalinas/StylelintBC/internal/constants"

func GetErrorMessage(error string) constants.ErrorMessage {
	return constants.ERROR_MESSAGES[error]
}
