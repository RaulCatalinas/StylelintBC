package utils

import "stylelintbc/src/constants"

func GetErrorMessage(error string) constants.ErrorMessage {
	return constants.ERROR_MESSAGES[error]
}
