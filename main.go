package main



func main() {

}


func CelciusToFahrenheit(suhuCelsius float64) float64 {
	suhuFahrenheit := ((9.0 / 5.0) * suhuCelsius) + 32
	return suhuFahrenheit
}

func CelciusToKelvin(suhuCelsius float64) float64 {
	suhuKelvin := suhuCelsius + 273.15
	return suhuKelvin
}

func FahrenheitToCelcius(suhuFahrenheit float64) float64 {
	suhuCelsius := (suhuFahrenheit - 32) * (5.0 / 9.0)
	return suhuCelsius
}

func FahrenheitToKelvin(suhuFahrenheit float64) float64 {
	suhuKelvin := (suhuFahrenheit + 459.67) * (5.0 / 9.0)
	return suhuKelvin
}

func KelvinToCelsius(suhuKelvin float64) float64 {
	suhuCelsius := suhuKelvin - 273.15
	return suhuCelsius
}

func KelvinToFahrenheit(suhuKelvin float64) float64 {
	suhuFahrenheit := (suhuKelvin * (9.0 / 5.0)) - 459.67
	return suhuFahrenheit
}