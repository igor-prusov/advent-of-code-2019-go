package main

import "testing"

func TestFuelForMass(t *testing.T) {
	cases := []struct {
		input, expected int64
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

func TestFuelForMassPrecise(t *testing.T) {
	cases := []struct {
		input, expected int64
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, c := range cases {
		got := fuelForMassPrecise(c.input)
		if got != c.expected {
			t.Errorf("fuelForMassPrecise(%v) == %v, want %v", c.input, got, c.expected)
		}

	}
}
