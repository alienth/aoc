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
const OpcodeJumpIfTrue = 5
const OpcodeJumpIfFalse = 6
const OpcodeLessThan = 7
const OpcodeEquals = 8
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

	software := make([]int, 0)
	s := bufio.NewScanner(f)
	s.Scan()
	line := s.Text()
	items := strings.Split(line, ",")
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		software = append(software, i)
	}
	phaseSettings := []int{5, 6, 7, 8, 9}

	// map of amplifiers
	amplifiers := make(map[int]*amplifier)

	largest := 0
	lastOutputSignal := 0
	inputSignal := 0
	for _, phaseArray := range permutations(phaseSettings) {
		halted := false

		for i, p := range phaseArray {
			ampMemory := make([]int, len(software))
			copy(ampMemory, software)
			amplifiers[i] = &amplifier{phase: p, software: ampMemory}
		}

		for !halted {
			for i := range phaseArray {
				inputSignal, halted = amplifiers[i].execute(inputSignal)
				if i == 4 && !halted {
					fmt.Println("Storing last output signal from E, ", inputSignal)
					lastOutputSignal = inputSignal
				}
			}
		}
		fmt.Println("halted! ", lastOutputSignal)
		if lastOutputSignal > largest {
			largest = lastOutputSignal
		}

	}
	fmt.Printf("%d\n", largest)
}

type amplifier struct {
	phase    int
	phaseSet bool
	software []int
	pos      int
}

func (a *amplifier) execute(inputSignal int) (int, bool) {
	fmt.Println(a.phase, inputSignal)

	output := 0
	for a.pos < len(a.software) {
		pos := a.pos
		// fmt.Println("At position ", pos)
		switch instruction := processInstruction(a.software[pos]); instruction.opcode {
		case OpcodeAdd, OpcodeMultiply, OpcodeLessThan, OpcodeEquals: // 3 params
			math(a.software, pos, instruction)
			a.pos += 4
		case OpcodeJumpIfTrue: // 2 params
			fmt.Printf("Instruction: %v\n", a.software[pos:pos+3])
			paramValues := getParamValues(a.software, instruction.paramModes, a.software[pos+1:pos+3])
			if paramValues[0] != 0 {
				a.pos = paramValues[1]
			} else {
				a.pos += 3
			}
			fmt.Printf("Jumping to: %v\n", a.pos)
		case OpcodeJumpIfFalse: // 2 params
			fmt.Printf("Instruction: %v\n", a.software[pos:pos+3])
			paramValues := getParamValues(a.software, instruction.paramModes, a.software[pos+1:pos+3])
			if paramValues[0] == 0 {
				a.pos = paramValues[1]
			} else {
				a.pos += 3
			}
			fmt.Printf("Jumping to: %v\n", a.pos)
		case OpcodeInput: // 1 param
			fmt.Printf("Instruction: %v\n", a.software[pos:pos+2])
			store := a.software[pos+1]
			if !a.phaseSet {
				a.software[store] = a.phase
				a.phaseSet = true
			} else {
				a.software[store] = inputSignal
			}
			a.pos += 2
		case OpcodeOutput: // 1 param
			fmt.Printf("Instruction: %v\n", a.software[pos:pos+2])
			paramValues := getParamValues(a.software, instruction.paramModes, []int{a.software[pos+1]})
			fmt.Println("Output: ", paramValues[0])
			output = paramValues[0]
			a.pos += 2
			return output, false
		case OpcodeEnd:
			return output, true
		default:
			log.Fatalf("Unknown opcode %d on instruction %d.\n", instruction.opcode, a.software[pos])
		}
	}
	return output, false
}

func math(input []int, pos int, ins instruction) {
	fmt.Printf("Instruction: %v\n", input[pos:pos+4])
	params := input[pos+1 : pos+3]
	paramValues := getParamValues(input, ins.paramModes, params)
	store := input[pos+3]
	result := 0

	switch ins.opcode {
	case OpcodeAdd:
		result = add(paramValues)
		fmt.Printf("Adding %v and storing sum %d at address %d.\n", paramValues, result, store)
	case OpcodeMultiply:
		result = multiply(paramValues)
		fmt.Printf("Multiplying %v and storing product %d at address %d.\n", paramValues, result, store)
	case OpcodeLessThan:
		if paramValues[0] < paramValues[1] {
			result = 1
		} else {
			result = 0
		}
		fmt.Printf("Less than on %v and storing result %d at address %d.\n", paramValues, result, store)
	case OpcodeEquals:
		if paramValues[0] == paramValues[1] {
			result = 1
		} else {
			result = 0
		}
		fmt.Printf("Equals on %v and storing result %d at address %d.\n", paramValues, result, store)
	}

	input[store] = result
}

func add(operands []int) int {
	sum := 0
	for _, o := range operands {
		sum += o
	}
	return sum
}

func multiply(operands []int) int {
	return operands[0] * operands[1]
}

type instruction struct {
	opcode     int
	paramModes []paramMode
}

func getParamValues(input []int, modes []paramMode, params []int) []int {
	// fmt.Println("Modes ", modes)
	// fmt.Println("Params ", params)
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
	// fmt.Println("Values ", values)
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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
