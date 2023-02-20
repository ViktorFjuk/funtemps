package conv

import (
	"reflect"
	"testing"
)

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

	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 25, want: 298.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}

func TestKelvinToFahrenheit(t *testing.T) {

	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 300, want: 80.33000000000004},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}

func TestKelvinToCelsius(t *testing.T) {

	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 300, want: 26.850000000000023},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}
