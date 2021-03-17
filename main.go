package main

import (
	"fmt"
	"strings"
	"temp/convert"
	"temp/input"
)

func main() {

	var temp float64
	var unit string

	input.TakeTemperature("temperature ?", &temp)
	input.TakeUnite("chose unit pres 'C' for celsious 'F' for fahrenheit ", &unit)

	if strings.ToUpper(strings.TrimSpace(unit)) == "C" {
		convert.ToCelsius(&temp, &unit)
	} else {
		convert.ToFahrenheit(&temp, &unit)
	}

	fmt.Println("temp in", unit, " :", temp)

}
