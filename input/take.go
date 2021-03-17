package input

import (
	"errors"
	"fmt"
)

var errReadingInput = errors.New("Error reading input")

func TakeTemperature(msg string, value *float64) {
	fmt.Println(msg)
	_, err := fmt.Scanln(value)
	if err != nil {
		Error(errReadingInput)
	}
}

func TakeUnite(msg string, value *string) {
	fmt.Println(msg)
	_, err := fmt.Scanln(value)
	if err != nil {
		Error(errReadingInput)
	}
}
