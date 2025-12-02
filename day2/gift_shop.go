package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// checkForValidIDs returns false if the ID is invalid (if they are a sequence of digits repeated twice)
// and otherwise returns true
func checkForValidID(id int) bool {
	digits := len(strconv.Itoa(id))

	// if there's an odd number of digits, it's flat-out valid
	if digits%2 == 1 {
		return true
	}

	// If the int is split in half and the first half equals the second half, it is invalid
	firstHalf := id / int(math.Pow(10, float64(digits/2)))
	secondHalf := id % int(math.Pow(10, float64(digits/2)))
	if firstHalf-secondHalf == 0 {
		return false
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
