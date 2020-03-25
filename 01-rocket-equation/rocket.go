package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuelForMass(mass uint64) uint64 {
	return mass/3 - 2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err.Error())
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)

	var fuelMass uint64 = 0

	for scanner.Scan() {
		mass, err := strconv.ParseUint(scanner.Text(), 10, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: ", err.Error())
			os.Exit(1)
		}
		fuelMass += fuelForMass(mass)
	}
	fmt.Println("Fuel required = ", fuelMass)
}
