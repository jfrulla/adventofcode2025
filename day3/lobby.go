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
	joltage := 0

	// Slurp up the input file
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// For each line in the input file
	lines := strings.Split(string(inputFileContents), "\n")
	for _, line := range lines {
		if line == "" {
			continue // skip empty lines
		}
		joltage += getBatteryBankJoltage(line)
	}
	fmt.Printf("The joltage is %d\n", joltage)
}

// getBatteryBankJoltage takes a string representing a battery bank
// expected format is just a series of integers: "867442382356..."
// It finds the greatest possible Joltage and returns that as a result
func getBatteryBankJoltage(bank string) int {
	type battery struct {
		index int
		value int
	}

	// I need to keep track of value AND position
	var highestFirstDigit = battery{index: 0, value: 0}
	var highestSecondDigit = battery{index: 1, value: 0}

	bankBatteries := len(bank)

	for index, rune := range bank {
		digit, _ := strconv.Atoi(string(rune))

		// If we find a battery that is higher than the first digit,
		// it doesn't matter what the second digit was before.
		if index >= highestFirstDigit.index && index < bankBatteries-1 {
			if digit > highestFirstDigit.value {
				highestFirstDigit.index = index
				highestFirstDigit.value = digit
				highestSecondDigit.index = index + 1
				highestSecondDigit.value = 0
				continue
				// if we found a new first battery,
				// there's no reason to look at the second battery
			}
		}
		if index >= highestSecondDigit.index {
			if digit > highestSecondDigit.value {
				highestSecondDigit.index = index
				highestSecondDigit.value = digit
			}
		}
	}

	// conversion BS to concatenate the two digits together
	joltage, _ := strconv.Atoi(fmt.Sprintf("%d%d", highestFirstDigit.value, highestSecondDigit.value))
	return joltage
}
