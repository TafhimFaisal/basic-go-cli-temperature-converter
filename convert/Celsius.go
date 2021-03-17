package convert

func ToCelsius(temp *float64, unitname *string) {
	*temp = (*temp - 32) * 5 / 9
	*unitname = "Celsius"
}
