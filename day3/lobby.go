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
		joltage += getBatteryBankJoltage(line, 12)
	}
	fmt.Printf("The joltage is %d\n", joltage)
}

// getBatteryBankJoltage takes a string representing a battery bank
// expected format is just a series of integers: "867442382356..."
// It finds the greatest possible Joltage and returns that as a result
func getBatteryBankJoltage(bank string, batteriesOn int) int {
	type battery struct {
		index int
		value int
	}
	activatedBatteries := make([]battery, batteriesOn)
	numBankBatteries := len(bank)

	for index, batRune := range bank {
		digit, _ := strconv.Atoi(string(batRune))

		// If we find a battery in the bank that needs to go in our slice of active batteries,
		// then set it and clear out the values of every battery after it in our
		// slice of activated batteries.
		clearRemaining := false
		for i, b := range activatedBatteries {
			if clearRemaining {
				activatedBatteries[i].index = 0
				activatedBatteries[i].value = 0
				continue
			}

			// e.g. the third-to-last battery in the bank CANNOT be
			// anything before our third-to-last activated battery
			finalIndex := numBankBatteries - batteriesOn + i
			if index >= b.index && index <= finalIndex && digit > b.value {
				activatedBatteries[i].index = index
				activatedBatteries[i].value = digit
				clearRemaining = true
			}
		}
	}

	// concatenate all digits together
	joltageStr := ""
	for _, bat := range activatedBatteries {
		joltageStr = fmt.Sprintf("%s%d", joltageStr, bat.value)
	}
	joltage, _ := strconv.Atoi(joltageStr)
	return joltage
}
