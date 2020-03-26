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

func programRun(program []int) int {
	for i := 0; ; i += 4 {
		switch opcode := program[i]; opcode {
		case 99:
			return 0
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
			return -1

		}
	}
}

func restoreProgram(program []int, noun int, verb int) {
	program[1] = noun
	program[2] = verb
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err.Error())
		os.Exit(1)
	}

	inpunt_string := strings.Trim(string(input), "\n")
	original_program := parseInput(inpunt_string)

	program := make([]int, len(original_program))
	copy(program, original_program)

	restoreProgram(program, 12, 2)
	programRun(program)

	fmt.Println("Result =", program[0])

	const desired_output = 19690720

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(program, original_program)

			restoreProgram(program, noun, verb)
			res := programRun(program)
			if res == 0 && program[0] == desired_output {
				fmt.Println("Result =", 100*noun+verb)
				return
			}
		}
	}
	os.Exit(1)
}
