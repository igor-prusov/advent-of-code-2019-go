package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuelForMass(mass int64) int64 {
	return mass/3 - 2
}

func fuelForMassPrecise(mass int64) int64 {
	if fuel := fuelForMass(mass); fuel < 0 {
		return 0
	} else {
		return fuel + fuelForMassPrecise(fuel)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err.Error())
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)

	var fuelMass int64 = 0
	var fuelMassPrecise int64 = 0

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: ", err.Error())
			os.Exit(1)
		}
		fuelMass += fuelForMass(int64(mass))
		fuelMassPrecise += fuelForMassPrecise(int64(mass))
	}
	fmt.Println("Fuel required (fuel mass == 0) = ", fuelMass)
	fmt.Println("Fuel required (fuel mass != 0) = ", fuelMassPrecise)
}
