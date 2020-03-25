package main

import (
	"fmt"
)

func fuelForMass(mass uint) uint {
	return mass/3 - 2
}

func main() {
	var totalFuel uint
	for {
		var mass uint
		_, err := fmt.Scanf("%d", &mass)
		if err != nil {
			break
		}
		totalFuel += fuelForMass(mass)
	}
	fmt.Println("Total fuel = ", totalFuel)
}
