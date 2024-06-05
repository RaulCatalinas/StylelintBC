package utils

import "stylelintbc/constants"

func GetErrorMessage(error string) constants.ErrorMessage {
	return constants.ERROR_MESSAGES[error]
}
