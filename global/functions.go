package global

import (
	"fmt"

	"github.com/bbruun/galt/initializers"
)

func Debug(text string) {
	if initializers.DEBUG {
		if text[len(text)-1:] == "\n" {
			fmt.Printf("DEBUG %s", text)
		} else {
			fmt.Printf("DEBUG %s\n", text)
		}
	}
}
