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
	operands, operators, leftAlignment, err := processInput(inputFileContents)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	//part 1
	total := cephMath(operands, operators)
	fmt.Printf("The part1 answer is %d\n", total)

	//part 2
	operands = transformElements(operands, leftAlignment)
	total = cephMath(operands, operators)
	fmt.Printf("The part2 answer is %d\n", total)
}

// procesInput takes the entire bytestring from the file and
// returns a list of operations ([]rune)
// and a list of operands per operation ([][]int)
func processInput(fileInput []byte) ([][]int, []rune, []bool, error) {
	// get our datastructures ready
	operands := [][]int{}
	intMatrix := [][]int{}
	operators := []rune{}
	leftAlignment := []bool{}

	// For each line in the input file
	lines := strings.Split(string(fileInput), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		_, err := strconv.Atoi(parts[0])
		if err != nil {
			// must be the line of operators
			// loop through each individual character,
			// get the operators, and also
			// determine whether the column is left or right aligned
			alignmentIndex := 0
			for charIndex, char := range line {
				if char != ' ' {
					leftAlignment = append(leftAlignment, true)
					operators = append(operators, char)
					// for each line
					for _, l := range lines {
						if l == "" {
							continue
						}
						// the operator is ALWAYS left aligned, so any spaces
						// in the same column mean the operands are right-aligned
						// (or square)
						if l[charIndex] == ' ' {
							leftAlignment[alignmentIndex] = false
						}
					}
					alignmentIndex++
				}
			}
		} else {
			// must be a number
			intSlice := []int{}
			for _, field := range parts {
				k, _ := strconv.Atoi(field)
				intSlice = append(intSlice, k)
			}
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

	return operands, operators, leftAlignment, nil
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

func transformElements(operands [][]int, leftAlignment []bool) [][]int {
	newOperands := [][]int{}
	for index, opList := range operands {
		newStrOps := []string{}
		newOps := []int{}
		for _, op := range opList {
			opStr := strconv.Itoa(op)
			width := len(opStr)
			if leftAlignment[index] {
				for i := 0; i < width; i++ {
					if len(newStrOps) < width {
						newStrOps = append(newStrOps, "")
					}
					newStrOps[i] += string(opStr[i])
				}
			} else {
				for i := width - 1; i >= 0; i-- {
					if len(newStrOps) < width {
						newStrOps = append(newStrOps, "")
					}
					newStrOps[width-1-i] += string(opStr[i])
				}
			}
		}
		for _, opStr := range newStrOps {
			newOperand, _ := strconv.Atoi(opStr)
			newOps = append(newOps, newOperand)
		}
		newOperands = append(newOperands, newOps)
	}
	return newOperands
}
