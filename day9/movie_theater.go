package main

import (
	"log"
	"os"
	"strings"
)

// MAIN
func main() {
	// Slurp up the input
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// For each line in the input file
	lines := strings.Split(string(inputFileContents), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		// iterate through the input
		// record the leftmost, rightmost, topmost, and bottommost points
		// the answer will PROBABLY be a square with a diagonal created
		// between the topleft&bottomright or the topright&bottomleft

		// might need to bound it by getting the leftmost and rightmost points of every line
		// but only if it is further left or right than previous lines

		// for every line, add up-to two points to our list of viable corners
		// if more left or more right, then drop previous viable points
		// ((that are no longer viable)) <- interesting algo needed to determine

		// should be able to draw a wonky circle by connecting
		// all our viable points with lines and have NO points fall outside of those lines
	}
}
