package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

func FahrenheitToCelsius(fahrenheit float64) float64 {

	return (fahrenheit - 32) * 5 / 9
}

func FahrenheitToKelvin(fahrenheit float64) float64 {

	return (fahrenheit + 459.67) * 5 / 9
}

func CelsiusToFahrenheit(celsius float64) float64 {

	return (celsius * 9 / 5) + 32
}

func CelsiusToKelvin(value float64) float64 {
	return 0
}

func KelvinToFahrenheit(value float64) float64 {
	return 0
}

func KelvinToCelsius(value float64) float64 {
	return 0
}
