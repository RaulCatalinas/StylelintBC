package cli

import (
	"fmt"

	"github.com/RaulCatalinas/StylelintBC/internal/constants"
)

func ShowVersion() {
	fmt.Println("StylelintBC -> " + constants.VERSION)
}
