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
func checkForValidID(id int) bool {
	idString := strconv.Itoa(id)
	digits := len(idString)
	for i := 1; i <= (digits / 2); i++ {
		if digits%i != 0 {
			continue // not an equal pattern, couldn't possibly be invalid
		}
		j := digits / i // might be a pattern of numbers repeating j times
		for k := 1; k < j; k++ {
			groupA := idString[(k-1)*i : k*i]
			groupB := idString[k*i : (k+1)*i]
			if groupA != groupB {
				break
			}
			if k == j-1 {
				return false
			}
		}
	}
	return true
}

// sumBadIDs takes one line of input as a string
// we expect the input to be a list of integer ranges
// e.g. "1234-4567,890-1023,987876-1234123445"
func sumBadIDs(allIdRanges string) int {
	sumOfBadIDs := 0
	idRanges := strings.Split(allIdRanges, ",")
	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		firstNum, _ := strconv.Atoi(parts[0])
		lastNum, _ := strconv.Atoi(parts[1])
		for i := firstNum; i <= lastNum; i++ {
			if !checkForValidID(i) {
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
