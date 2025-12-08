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
	junctionBoxCoords := [][]int{} // it'll be a list of {x, y, z} coordinates
	answer := 0

	// For each line in the input file
	lines := strings.Split(string(inputFileContents), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		vector := []int{}
		parts := strings.Split(strings.TrimSpace(line), ",")
		for _, posStr := range parts {
			position, _ := strconv.Atoi(posStr)
			vector = append(vector, position)
		}
		junctionBoxCoords = append(junctionBoxCoords, vector)
	}

	//part 1
	fmt.Printf("The part1 answer is %d\n", answer)
}
