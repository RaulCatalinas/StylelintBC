package cli

import (
	"fmt"

	"github.com/RaulCatalinas/stylelintbc/internal/constants"
)

func ShowVersion() {
	fmt.Println("StylelintBC -> " + constants.VERSION)
}
