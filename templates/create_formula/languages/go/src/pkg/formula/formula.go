package formula

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

type Formula struct {
	Text    string
	List    string
	Boolean string
}

func (h Formula) Run(writer io.Writer) {
	var result string
	result += fmt.Sprintf("Hello world!\n")
	result += color.FgGreen.Render(fmt.Sprintf("You receive %s in text.\n", h.Text))
	result += color.FgRed.Render(fmt.Sprintf("You receive %s in list.\n", h.List))
	result += color.FgYellow.Render(fmt.Sprintf("You receive %s in boolean.\n", h.Boolean))

	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}
