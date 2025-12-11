package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	lightDiagram uint64
	buttons      []uint64
	joltages     []int
}

// MAIN
func main() {
	// Slurp up the input file
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// process & parse the input file
	machines, err := processInput(inputFileContents)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	//part 1
	total := sumMachinePresses(machines)
	fmt.Printf("The part1 answer is %d\n", total)
}

// procesInput takes the entire bytestring from the file
func processInput(fileInput []byte) ([]Machine, error) {
	machines := []Machine{}

	// For each line in the input file
	lines := strings.Split(string(fileInput), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var inputLights uint64
		inputButtonList := []uint64{}
		inputJoltages := []int{}

		// for each space-separated field
		parts := strings.Fields(line)
		binaryString := ""
		for _, part := range parts {

			// this part is an indicator light diagram
			if part[0] == '[' {
				for i := 1; i < (len(part) - 1); i++ {
					if part[i] == '.' {
						binaryString += "0"
					} else {
						binaryString += "1"
					}
				}
				inputLights, _ = strconv.ParseUint(binaryString, 2, 0)
			}

			// this part is a button wiring schematic
			if part[0] == '(' {
				byteStr := []byte{}
				for j := 0; j < len(binaryString); j++ {
					byteStr = append(byteStr, '0')
				}
				for i := 1; i < (len(part) - 1); i++ {
					num, err := strconv.Atoi(string(part[i]))
					if err != nil {
						continue
					}
					byteStr[num] = '1'
				}
				inputButtonUint, _ := strconv.ParseUint(string(byteStr), 2, 0)
				inputButtonList = append(inputButtonList, inputButtonUint)
			}

			// this part is a button wiring schematic
			if part[0] == '{' {
				for i := 1; i < (len(part) - 1); i++ {
					num, _ := strconv.Atoi(string(part[i]))
					inputJoltages = append(inputJoltages, num)
				}
			}
		}
		machines = append(machines,
			Machine{
				lightDiagram: inputLights,
				buttons:      inputButtonList,
				joltages:     inputJoltages,
			})
	}
	return machines, nil
}

func sumMachinePresses(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		presses := 0
		if machine.lightDiagram == uint64(0) {
			continue
		}
		for i := 1; i <= 20; i++ {
			if getButtonPresses(machine.lightDiagram, machine.buttons, i) {
				total += i
				presses = i
				break
			}
			if i == 20 {
				presses = 1000
			}
		}
		fmt.Printf("presses for this machine: %d\n", presses)
	}
	return total
}

func getButtonPresses(bsInput uint64, bsList []uint64, depth int) bool {
	for _, bs := range bsList {
		if bsInput^bs == uint64(0) {
			return true
		}
	}
	for _, bs := range bsList {
		if depth == 1 {
			return false
		}
		if getButtonPresses(bsInput^bs, bsList, depth-1) {
			return true
		}
	}
	return false
}
