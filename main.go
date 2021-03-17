package main

import (
	"fmt"
	"strings"
	"temp/convert"
)

func main() {
	var temp float64
	var unit string

	fmt.Println("temperature ?")
	fmt.Scanln(&temp)

	fmt.Println("chose unit pres 'C' for celsious 'F' for fahrenheit ")
	fmt.Scanln(&unit)

	if strings.ToUpper(strings.TrimSpace(unit)) == "C" {
		convert.ToCelsius(&temp, &unit)
	} else {
		convert.ToFahrenheit(&temp, &unit)
	}

	fmt.Println("temp in", unit, " :", temp)

}
