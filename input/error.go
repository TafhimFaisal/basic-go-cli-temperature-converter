package input

import (
	"fmt"
	"os"
)

func Error(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
