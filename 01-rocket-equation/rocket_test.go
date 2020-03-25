package main

import "testing"

func TestFuelForMass(t *testing.T) {
	cases := []struct {
		input, expected uint
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, c := range cases {
		got := fuelForMass(c.input)
		if got != c.expected {
			t.Errorf("fuelForMass(%v) == %v, want %v", c.input, got, c.expected)
		}

	}
}
