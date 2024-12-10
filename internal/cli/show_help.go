package cli

import (
	"fmt"
	"strings"

	"github.com/RaulCatalinas/stylelintbc/internal/types"
)

func ShowHelp(options []types.Option) {
	var builder strings.Builder

	builder.WriteString("Usage: stylelintbc [options]\n\n")
	builder.WriteString("Command line for easy Stylelint configuration\n\n")
	builder.WriteString("Options:\n")

	for _, option := range options {
		fmt.Fprintf(
			&builder,
			"%-15s %-5s %s\n",
			option.Name,
			option.Alias,
			option.Description,
		)
	}

	builder.WriteString("\n")

	fmt.Print(builder.String())
}
