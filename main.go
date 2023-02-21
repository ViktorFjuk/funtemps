package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/ViktorFjuk/funtemps/conv"
)

var fahr float64
var cel float64
var kel float64
var out string

func init() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")


}

func main() {
	
	flag.Parse()

