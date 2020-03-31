package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
)

func usage() {
	log.Fatal("Usage:", path.Base(os.Args[0]), " <range>")
}

func parseInput(input string) (int, int, error) {
	var start int
	var end int
	_, err := fmt.Sscanf(input, "%d-%d\n", &start, &end)
	if err != nil {
		start, end = 0, 0
	}
	return start, end, err
}

func verifyRange(start int, end int) error {
	switch {
	case start > end:
		return errors.New("start < end")
	case start < 100000:
		return errors.New("interval start should be a 6-digit number")
	case end > 999999:
		return errors.New("interval end should be a 6-digit number")
	}
	return nil
}

func verifyPassword(password int) (bool, bool) {
	var hasStrictlyDouble bool
	var hasDouble bool
	repeatCounter := 0
	prev := 10
	d, password := password%10, password/10
	for d > 0 || password > 0 {
		if d > prev {
			return false, false
		}
		if d == prev {
			repeatCounter++
			hasDouble = true
		} else {
			if repeatCounter == 1 {
				hasStrictlyDouble = true
			}
			repeatCounter = 0
		}
		prev, d, password = d, password%10, password/10

	}
	if repeatCounter == 1 {
		hasStrictlyDouble = true
		hasDouble = true
	}
	return hasDouble, hasStrictlyDouble
}

func bruteforce(start int, end int) (int, int) {
	var counter int
	var strictCounter int

	for current := start; current <= end; current++ {
		soft, strict := verifyPassword(current)
		if strict {
			strictCounter++
		}
		if soft {
			counter++
		}

	}
	return counter, strictCounter
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	input := os.Args[1]

	start, end, err := parseInput(input)
	if err != nil {
		log.Fatal("Incorrect range: ", input, ": ", err.Error())
	}

	if err := verifyRange(start, end); err != nil {
		log.Fatal("Invalid range: ", start, "-", end, "(", err.Error(), ")")
	}

	fmt.Println("range: ", start, "-", end)

	result1, result2 := bruteforce(start, end)
	fmt.Println(result1)
	fmt.Println(result2)
}
