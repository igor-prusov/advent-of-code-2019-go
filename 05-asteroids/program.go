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

func paresInstruction(input int) (instruction, error) {
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
			return inst, fmt.Errorf("Incorrect operand type")
		}
		input /= 10
	}
	return inst, nil
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
		inst, err := paresInstruction(program[i])
		if err != nil {
			return 0, nil
		}

		getOperand := func(j int) (int, error) {
			if inst.operands[j] == positionMode {
				return program[program[i+j+1]], nil
			} else if inst.operands[j] == immediateMode {
				return program[i+j+1], nil
			} else {
				return 0, fmt.Errorf("Incorrect operand type")
			}
		}

		switch inst.opcode {
		case 99:
			return 0, nil
		case 1:
			operand1, err := getOperand(0)
			if err != nil {
				return 0, err
			}

			operand2, err := getOperand(1)
			if err != nil {
				return 0, err
			}

			if inst.operands[2] == immediateMode {
				return 0, fmt.Errorf("Destination operand must be immediate")
			}

			resultIndex := program[i+3]
			program[resultIndex] = operand1 + operand2
			i += 4
		case 2:
			operand1, err := getOperand(0)
			if err != nil {
				return 0, err
			}

			operand2, err := getOperand(1)
			if err != nil {
				return 0, err
			}

			if inst.operands[2] == immediateMode {
				return 0, fmt.Errorf("Destination operand must be immediate")
			}

			resultIndex := program[i+3]
			program[resultIndex] = operand1 * operand2
			i += 4
		case 3:
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
