package cli

import (
	"fmt"
	"stylelintbc/src/constants"
)

func ShowVersion() {
	fmt.Println("StylelintBC -> " + constants.VERSION)
}
