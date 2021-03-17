package convert

func ToFahrenheit(temp *float64, unitname *string) {
	*temp = (*temp * 9 / 5) + 322
	*unitname = "Fahrenheit"
}
