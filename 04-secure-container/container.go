package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func usage() {
	log.Fatal("Usage:", path.Base(os.Args[0]), " <range>")
}

func parseInput(input string) (int, int) {
	var start int
	var end int
	_, err := fmt.Sscanf(input, "%d-%d", &start, &end)
	if err != nil {
		log.Fatal("Incorrect range: ", input)
	}
	return start, end
}

func verifyRange(start int, end int) bool {
	switch {
	case start > end:
		return false
	}
	return true
}

func verifyPassword(password int) bool {
	var hasDouble bool
	prev := 10
	d, password := password%10, password/10
	for d > 0 || password > 0 {
		if d > prev {
			return false
		}
		if d == prev {
			hasDouble = true
		}
		prev, d, password = d, password%10, password/10

	}
	return hasDouble
}

func bruteforce(start int, end int) int {
	var counter int

	for current := start; current <= end; current++ {
		if verifyPassword(current) {
			counter++
		}

	}
	return counter
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	start, end := parseInput(os.Args[1])
	if !verifyRange(start, end) {
		log.Fatal("Invalid range: ", start, "-", end)
	}

	fmt.Println("range: ", start, "-", end)

	result := bruteforce(start, end)
	fmt.Println(result)
}
