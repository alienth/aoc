package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const OpcodeAdd = 1
const OpcodeMultiply = 2
const OpcodeInput = 3
const OpcodeOutput = 4
const OpcodeEnd = 99

type paramMode int

const (
	paramModePosition paramMode = iota
	paramModeImmediate
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make([]int, 0)
	s := bufio.NewScanner(f)
	s.Scan()
	line := s.Text()
	items := strings.Split(line, ",")
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}
	execute(input)

	// for x := 0; x <= 99; x += 1 {
	// 	for y := 0; y <= 99; y += 1 {
	// 		m := make([]int, len(input))
	// 		copy(m, input)
	// 		m[1] = x
	// 		m[2] = y
	// 		result := execute(m)
	// 		if result == 19690720 {
	// 			fmt.Println(x, y)
	// 			fmt.Println((100 * x) + y)
	// 			return
	// 		}
	// 	}
	// }
}

func execute(input []int) int {
	last := len(input) - 1

	for pos := 0; pos <= last; {
		// fmt.Println("At position ", pos)
		switch instruction := processInstruction(input[pos]); instruction.opcode {
		case OpcodeAdd: // 3 params
			fmt.Printf("Instruction: %v\n", input[pos:pos+4])
			params := input[pos+1 : pos+3]
			paramValues := getParamValues(input, instruction.paramModes, params)
			operands := paramValues[:len(paramValues)]
			store := input[pos+3]
			result := add(operands)
			fmt.Printf("Adding %v and storing sum %d at address %d.\n", operands, result, store)
			input[store] = result
			pos += 4
		case OpcodeMultiply: // 3 params
			fmt.Printf("Instruction: %v\n", input[pos:pos+4])
			params := input[pos+1 : pos+3]
			paramValues := getParamValues(input, instruction.paramModes, params)
			operands := paramValues[:len(paramValues)]
			store := input[pos+3]
			result := multiply(operands)
			fmt.Printf("Multiplying %v and storing product %d at address %d.\n", operands, result, store)
			input[store] = result
			pos += 4
		case OpcodeInput:
			fmt.Printf("Instruction: %v\n", input[pos:pos+2])
			store := input[pos+1]
			input[store] = 1 // Our "input"
			pos += 2
		case OpcodeOutput:
			fmt.Printf("Instruction: %v\n", input[pos:pos+2])
			fmt.Println("Output: ", input[input[pos+1]])
			pos += 2
		case OpcodeEnd:
			return input[0]
		default:
			log.Fatalf("Unknown opcode %d on instruction %d.\n", instruction.opcode, input[pos])
		}
	}
	return 0
}

func add(operands []int) int {
	sum := 0
	for _, o := range operands {
		sum += o
	}
	return sum
}

func multiply(operands []int) int {
	product := 0
	for _, o := range operands {
		product = product * o
	}
	return product
}

type instruction struct {
	opcode     int
	paramModes []paramMode
}

func getParamValues(input []int, modes []paramMode, params []int) []int {
	fmt.Println("Modes ", modes)
	fmt.Println("Params ", params)
	values := make([]int, len(params))

	for i, p := range params {
		mode := paramModePosition
		if len(modes) > i {
			mode = modes[i]
		}
		switch mode {
		case paramModePosition:
			values[i] = input[p]
		case paramModeImmediate:
			values[i] = p
		default:
			log.Fatal("Unknown param mode: ", mode)
		}
	}
	fmt.Println("Values ", values)
	return values
}

func processInstruction(i int) instruction {
	digits := getDigits(i)
	reverseSlice(digits)

	var ins instruction
	ins.paramModes = make([]paramMode, 0)
	ins.opcode = digits[0]
	if len(digits) > 1 {
		ins.opcode += digits[1] * 10
	}

	for i := 2; i < len(digits); i++ {
		ins.paramModes = append(ins.paramModes, paramMode(digits[i]))
	}

	return ins
}

func reverseSlice(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

// Takes an integer and splits it into its composite digits.
func getDigits(i int) []int {
	digitChars := strings.Split(fmt.Sprintf("%d", i), "")
	digits := make([]int, len(digitChars))

	for i, d := range digitChars {
		digits[i], _ = strconv.Atoi(d)
	}
	return digits
}
