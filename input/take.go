package input

import (
	"errors"
	"fmt"
)

var errReadingInput = errors.New("Error reading input")

func TakeFloat(msg string, value *float64) {
	fmt.Println(msg)
	_, err := fmt.Scanln(value)
	if err != nil {
		Error(errReadingInput)
	}
}

func TakeString(msg string, value *string) {
	fmt.Println(msg)
	_, err := fmt.Scanln(value)
	if err != nil {
		Error(errReadingInput)
	}
}
