package main

import (
	"errors"
	"fmt"
	"strings"
	"temp/convert"
	"temp/input"
)

func main() {

	var temp float64
	var unit string
	var takeagain string

	for {

		input.TakeFloat("temperature ?", &temp)
		input.TakeString("chose unit pres 'C' for celsious 'F' for fahrenheit ", &unit)

		switch unitName := strings.ToUpper(strings.TrimSpace(unit)); unitName {
		case "C":
			convert.ToCelsius(&temp, &unit)
			fmt.Println("temp in", unit, " :", temp)
		case "F":
			convert.ToFahrenheit(&temp, &unit)
			fmt.Println("temp in", unit, " :", temp)
		default:
			input.Error(errors.New("Sorry wrong unit"))
		}

		input.TakeString("Would you like to convert another temperature ? (y/n) ", &takeagain)
		if strings.ToUpper(strings.TrimSpace(takeagain)) != "Y" {
			fmt.Println("thank you.")
			break
		}

	}

}
