package cli

import (
	"fmt"
	"stylelintbc/constants"
)

func ShowVersion() {
	fmt.Println("StylelintBC -> " + constants.VERSION)
}
