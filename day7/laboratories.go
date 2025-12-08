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

	// get our datastructures ready
	manifoldDiagram := [][]rune{}
	splits := 0
	timelines := 0

	// For each line in the input file
	lines := strings.Split(string(inputFileContents), "\n")
	for lineNum, line := range lines {
		if line == "" {
			continue
		}
		manifoldDiagram = append(manifoldDiagram, []rune(line))
		if lineNum == 0 {
			continue
		}
		splitsThisLine := 0
		for charIndex, char := range line {
			if char == '.' && (manifoldDiagram[lineNum-1][charIndex] == 'S' || manifoldDiagram[lineNum-1][charIndex] == '|') {
				manifoldDiagram[lineNum][charIndex] = '|'
			}
			if char == '^' && manifoldDiagram[lineNum-1][charIndex] == '|' {
				manifoldDiagram[lineNum][charIndex-1] = '|'
				manifoldDiagram[lineNum][charIndex+1] = '|'
				splits++
				splitsThisLine++
				//timelines = timelines + 2
			}
		}
		timelines = timelines * splitsThisLine
	}

	for _, line := range manifoldDiagram {
		fmt.Printf("%s\n", string(line))
	}

	//part 1
	fmt.Printf("The part1 answer is %d\n", splits)

	//part 2
	fmt.Printf("The part2 answer is %d\n", timelines)

}

//part 2 .. we start with one world and a single tachyon pulse
// the tachyon pulse hits a splitter and now there are 2 worlds and 2 tachyon beams
// each of those tachyon beams hits a splitter
// there are now 3 tachyon beams and 4 worlds -- 2 of the beams merged into the same one
// 3 tachyon beams all hit 3 splitters
// now 4 tachyon beams (2 converged) and 8 worlds
// new worlds = worlds + (splits * 2 - 1)

// procesInput takes the entire bytestring from the file and
// returns a list of operations ([]rune)
// and a list of operands per operation ([][]int)
func processInput(fileInput []byte) ([]string, error) {
	// get our datastructures ready
	manifoldDiagram := []string{}

	// For each line in the input file
	lines := strings.Split(string(fileInput), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
	}

	return manifoldDiagram, nil
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
