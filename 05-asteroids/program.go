package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type operandMode int

const (
	positionMode operandMode = iota
	immediateMode
)

type instruction struct {
	operands [3]operandMode
	opcode   int
}

func paresInstruction(input int) instruction {
	var inst instruction
	inst.opcode = input % 100
	input /= 100
	for i := range inst.operands {
		mode := input % 10
		switch mode {
		case 0:
			inst.operands[i] = positionMode
		case 1:
			inst.operands[i] = immediateMode
		default:
			panic("Incorrect operand type")
		}
		input /= 10
	}
	return inst
}

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
		fmt.Println("i = ", i)
		inst := paresInstruction(program[i])
		fmt.Println("instruction = ", inst)
		switch inst.opcode {
		case 99:
			return 0, nil
		case 1:
			fmt.Printf("[%d %d %d %d]\n", program[i], program[i+1], program[i+2], program[i+3])
			var operand1 int
			var operand2 int

			if inst.operands[0] == positionMode {
				operand1 = program[program[i+1]]
			} else {
				operand1 = program[i+1]
			}

			if inst.operands[1] == positionMode {
				operand2 = program[program[i+2]]
			} else {
				operand2 = program[i+2]
			}

			if inst.operands[2] == immediateMode {
				panic("Unexpected")
			}

			resultIndex := program[i+3]
			program[resultIndex] = operand1 + operand2
			i += 4
		case 2:
			fmt.Printf("[%d %d %d %d]\n", program[i], program[i+1], program[i+2], program[i+3])
			var operand1 int
			var operand2 int

			if inst.operands[0] == positionMode {
				operand1 = program[program[i+1]]
			} else {
				operand1 = program[i+1]
			}

			if inst.operands[1] == positionMode {
				operand2 = program[program[i+2]]
			} else {
				operand2 = program[i+2]
			}

			if inst.operands[2] == immediateMode {
				panic("Unexpected")
			}

			resultIndex := program[i+3]
			program[resultIndex] = operand1 * operand2
			i += 4
		case 3:
			fmt.Printf("[%d %d]\n", program[i], program[i+1])
			var input int
			fmt.Print("Enter integer: ")
			_, err := fmt.Scanf("%d", &input)
			if err != nil {
				return 0, fmt.Errorf("Failed to parse user input")
			}
			address := program[i+1]
			program[address] = input
			i += 2
		case 4:
			fmt.Printf("[%d %d]\n", program[i], program[i+1])
			address := program[i+1]
			fmt.Println("Print: ", program[address])
			i += 2

		default:
			return 0, fmt.Errorf("Failed to parse opcode: %d", inst.opcode)

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

	inputString := strings.Trim(string(input), "\n")
	originalProgram := parseInput(inputString)

	program := make([]int, len(originalProgram))
	copy(program, originalProgram)

	// restoreProgram(program, 12, 2)
	_, err = programRun(program)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when executing user program:", err.Error())
		os.Exit(1)
	}
}
