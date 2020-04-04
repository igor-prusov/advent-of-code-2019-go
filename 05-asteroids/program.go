package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	programText := strings.Split(input, ",")
	var program = make([]int, len(programText))
	for i, cmd := range programText {
		val, err := strconv.Atoi(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error while parsing program: ", err.Error())
			os.Exit(1)
		}
		program[i] = val
	}
	return program
}

func programRun(program []int) (int, error) {
	for i := 0; ; {
		switch opcode := program[i]; opcode {
		case 99:
			return 0, nil
		case 1:
			operand1 := program[program[i+1]]
			operand2 := program[program[i+2]]
			resultIndex := program[i+3]
			program[resultIndex] = operand1 + operand2
			i += 4
		case 2:
			operand1 := program[program[i+1]]
			operand2 := program[program[i+2]]
			resultIndex := program[i+3]
			program[resultIndex] = operand1 * operand2
			i += 4
		case 3:
			var input int
			_, err := fmt.Scanf("%d", &input)
			if err != nil {
				return 0, fmt.Errorf("Failed to parse user input")
			}
			address := program[i+1]
			program[address] = input
			i += 2
		case 4:
			address := program[i+1]
			fmt.Println("Print: ", program[address])
			i += 2

		default:
			return 0, fmt.Errorf("Failed to parse opcode: %d", opcode)

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

	inpuntString := strings.Trim(string(input), "\n")
	originalProgram := parseInput(inpuntString)

	program := make([]int, len(originalProgram))
	copy(program, originalProgram)

	restoreProgram(program, 12, 2)
	programRun(program)

	fmt.Println("Result =", program[0])

	const desiredOutput = 19690720

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(program, originalProgram)

			restoreProgram(program, noun, verb)
			_, err := programRun(program)
			if err == nil && program[0] == desiredOutput {
				fmt.Println("Result =", 100*noun+verb)
				return
			}
		}
	}
	os.Exit(1)
}
