package main

import (
	"flag"
	"fmt"
	"github.com/ViktorFjuk/funtemps/conv"
)

var fahr float64
var cel float64
var kel float64
var out string

func init() {

}

func main() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

	flag.Parse()

	if cel != 0 {
		fmt.Printf("%.2fC = %.2fF = %.2fK\n", cel, conv.CelsiusToFahrenheit(cel), conv.CelsiusToKelvin(cel))
	} else if fahr != 0 {
		fmt.Printf("%.2fF = %.2fC = %.2fK\n", fahr, conv.FahrenheitToCelsius(fahr), conv.FahrenheitToKelvin(fahr))
	} else if kel != 0 {
		fmt.Printf("%.2fK = %.2fC = %.2fF\n", kel, conv.KelvinToCelsius(kel), conv.KelvinToFahrenheit(kel))
	} else {
		fmt.Println("Please specify a temperature to convert")
	}

}
