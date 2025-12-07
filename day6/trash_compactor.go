package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// MAIN
func main() {
	// Slurp up the input file
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// process & parse the input file
	operands, operators, err := processInputForPart1(inputFileContents)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	total := cephMath(operands, operators)
	fmt.Printf("The part1 answer is %d\n", total)

	//part 2
	operands, operators, err = processInputForPart1(inputFileContents)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	total = cephMath(operands, operators)
	fmt.Printf("The part2 answer is %d\n", total)

}

// procesInput takes the entire bytestring from the file and
// returns a list of operations ([]rune)
// and a list of operands per operation ([][]int)
func processInputForPart1(fileInput []byte) ([][]int, []rune, error) {
	// get our datastructures ready
	operands := [][]int{}
	intMatrix := [][]int{}
	operators := []rune{}

	// For each line in the input file
	lines := strings.Split(string(fileInput), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		intSlice := []int{}
		parts := strings.Fields(line)
		for _, field := range parts {
			k, err := strconv.Atoi(field)
			if err != nil {
				// must be a * or +
				operators = append(operators, []rune(field)[0])
			} else {
				intSlice = append(intSlice, k)
			}
		}
		if len(intSlice) > 0 {
			intMatrix = append(intMatrix, intSlice)
		}
	}

	// now we need to transform the intMatrix to operationElements
	// by reversing our rows and columns
	for index, _ := range operators {
		elementSlice := []int{}
		for _, intSlice := range intMatrix {
			elementSlice = append(elementSlice, intSlice[index])
		}
		operands = append(operands, elementSlice)
	}

	return operands, operators, nil
}

// procesInput takes the entire bytestring from the file and
// returns a list of operations ([]rune)
// and a list of operands per operation ([][]int)
func processInputForPart2(fileInput []byte) ([][]int, []rune, error) {
	// get our datastructures ready
	operands := [][]int{}
	intMatrix := [][]int{}
	operators := []rune{}

	// For each line in the input file
	lines := strings.Split(string(fileInput), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		intSlice := []int{}
		parts := strings.Fields(line)
		for _, field := range parts {
			k, err := strconv.Atoi(field)
			if err != nil {
				// must be a * or +
				operators = append(operators, []rune(field)[0])
			} else {
				intSlice = append(intSlice, k)
			}
		}
		if len(intSlice) > 0 {
			intMatrix = append(intMatrix, intSlice)
		}
	}

	// now we need to transform the intMatrix to operationElements
	// by reversing our rows and columns
	for index, _ := range operators {
		elementSlice := []int{}
		for _, intSlice := range intMatrix {
			elementSlice = append(elementSlice, intSlice[index])
		}
		operands = append(operands, elementSlice)
	}

	return operands, operators, nil
}

func cephMath(operands [][]int, operators []rune) int {
	sumSolutions := 0
	for i, op := range operators {
		solution := 0
		for j, operand := range operands[i] {
			if j == 0 {
				solution = operand
				continue
			}
			switch op {
			case '+':
				solution = solution + operand
				continue
			case '*':
				solution = solution * operand
				continue
			default:
				continue
			}
		}
		sumSolutions += solution
	}
	return sumSolutions
}

func transformElements(operands [][]int) [][]int {
	// iterate through the operands for each operation
	// convert the operands to strings
	// each number is going to be all operandString[i] concatenated together for each operand
	newOperands := [][]int{}
	for _, opList := range operands {
		newStrOps := []string{}
		newOps := []int{}
		for _, op := range opList {
			opStr := strconv.Itoa(op)
			width := len(opStr)
			// for each char in the opStr
			for i := width - 1; i >= 0; i-- {
				if len(newStrOps) < width {
					newStrOps = append(newStrOps, "")
				}
				newStrOps[width-1-i] += string(opStr[i])
			}
		}
		for _, opStr := range newStrOps {
			newOperand, _ := strconv.Atoi(opStr)
			newOps = append(newOps, newOperand)
		}
		newOperands = append(newOperands, newOps)
	}

	fmt.Printf("New Operands: %v\n", total)

	return newOperands
}
