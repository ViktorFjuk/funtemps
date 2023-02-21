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

func main() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

	flag.Parse()

	if cel != 0 {
		if out == "F" {
			fmt.Printf("%.2fC er %.2fF\n", cel, conv.CelsiusToFahrenheit(cel))
		} else if out == "K" {
			fmt.Printf("%.2fC er %.2fK\n", cel, conv.CelsiusToKelvin(cel))
		} else {
			fmt.Printf("%.2fC er %.2fC\n", cel, cel)
		}
	}

	if fahr != 0 {
		if out == "C" {
			fmt.Printf("%.2fF er %.2fC\n", fahr, conv.FahrenheitToCelsius(fahr))
		} else if out == "K" {
			fmt.Printf("%.2fF er %.2fK\n", fahr, conv.FahrenheitToKelvin(fahr))
		} else {
			fmt.Printf("%.2fF er %.2fF\n", fahr, fahr)
		}
	}

	if kel != 0 {
		if out == "C" {
			fmt.Printf("%.2fK er %.2fC\n", kel, conv.KelvinToCelsius(kel))
		} else if out == "F" {
			fmt.Printf("%.2fK er %.2fF\n", kel, conv.KelvinToFahrenheit(kel))
		} else {
			fmt.Printf("%.2fK er %.2fK\n", kel, kel)
		}
	}

}
