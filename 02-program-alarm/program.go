package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	program_text := strings.Split(input, ",")
	var program = make([]int, len(program_text))
	for i, cmd := range program_text {
		val, err := strconv.Atoi(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error while parsing program: ", err.Error())
			os.Exit(1)
		}
		program[i] = val
	}
	return program
}

func programRun(program []int) {
	for i := 0; ; i += 4 {
		switch opcode := program[i]; opcode {
		case 99:
			return
		case 1:
			operand1 := program[program[i+1]]
			operand2 := program[program[i+2]]
			result_index := program[i+3]
			program[result_index] = operand1 + operand2
		case 2:
			operand1 := program[program[i+1]]
			operand2 := program[program[i+2]]
			result_index := program[i+3]
			program[result_index] = operand1 * operand2
		default:
			fmt.Fprintln(os.Stderr, "Failed to parse opcode: ", opcode)
			os.Exit(1)

		}
	}
}

func restoreProgram(program []int) {
	program[1] = 12
	program[2] = 2
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err.Error())
		os.Exit(1)
	}

	inpunt_string := strings.Trim(string(input), "\n")
	program := parseInput(inpunt_string)
	restoreProgram(program)

	programRun(program)

	fmt.Println("Result =", program[0])
}
