package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// checkForValidIDs returns false if the ID is invalid (if they are a sequence of repeated numbers)
// and otherwise returns true
func isValidID(id int) bool {
	idString := strconv.Itoa(id)
	totalDigits := len(idString)

	// check every possible grouping of digits
	// starting at many groups of 1 digit, going to a max of 2 groups of half the digits each
	for sequenceDigits := 1; sequenceDigits <= (totalDigits / 2); sequenceDigits++ {
		if totalDigits%sequenceDigits != 0 {
			continue // not an equal pattern, try a different grouping
		}

		numGroups := totalDigits / sequenceDigits

		// iterate through pairs of digit groups
		// if any pair doesn't match, try again with a new sequenceDigits
		// if all pairs match, we've found a winner!
		for group := 1; group < numGroups; group++ {
			groupA := idString[(group-1)*sequenceDigits : group*sequenceDigits]
			groupB := idString[group*sequenceDigits : (group+1)*sequenceDigits]
			if groupA != groupB {
				break
			}
			if group == numGroups-1 {
				return false
			}
		}
	}
	return true
}

// sumBadIDs takes one line of input as a string
// we expect the input to be a list of integer ranges.
// e.g. "1234-4567,890-1023,987876-1234123445".
// returns the sum of all invalid IDs in the provided ranges.
func sumBadIDs(allIdRanges string) int {
	sumOfBadIDs := 0
	idRanges := strings.Split(allIdRanges, ",")
	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		firstNum, _ := strconv.Atoi(parts[0])
		lastNum, _ := strconv.Atoi(parts[1])
		for i := firstNum; i <= lastNum; i++ {
			if !isValidID(i) {
				sumOfBadIDs += i
			}
		}
	}
	return sumOfBadIDs
}

// MAIN
func main() {
	password := 0

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
		password += sumBadIDs(line)
	}
	fmt.Printf("The password is %d\n", password)
}
