package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// MAIN
func main() {

	totalFresh := 0

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
		// per line code here
	}

	fmt.Printf("The part1 answer is %d\n", totalFresh)
}

//func processInputLine
// need to separate fresh from spoiled ingredients
// given DB as puzzle input
// DB has ingredientIDs with fresh ranges, a blank line, and then all IDs
// all I need to do is get the list of IIDs and check to see if each ID
// is in any range
//
// Dude. This is just a binary search
