package conv

import (
	"reflect"
	"testing"
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 100, want: 37.77777777777778},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {

	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 90, want: 305.37222222222226},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}

func TestCelsiusToFahrenheit(t *testing.T) {

	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 25, want: 77},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}

func TestCelsiusToKelvin(t *testing.T) {

}

func TestKelvinToFahrenheit(t *testing.T) {
}

func TestKelvinToCelsius(t *testing.T) {

}
