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

	// this is (mostly) a minimum spanning tree problem.

	// Let each box be identified as its index in the junctionBoxCoords slice

	// Calculate all distances from each box to all other boxes.
	// That will be a O(n^2) calculation and it will be an n^2 sized list
	// in this case, thats a list of 1 million distances.
	// That's feasible for this input size, but that scales terribly.

	// hmm, it's not directional, so maybe only nlogn or something in that realm of complexity

	// sort box-pairs by distance and start connecting circuits
	// every time one circuit gets connected to another, the circuit itself becomes the box that it contains in the
	// list of distances and pairs

	// the circuits can be represented by a map[int]int which is map[boxIndex]=circuitIndex
	// initial values: circuit indexes and box indexes are the same.
	// whenever two boxes get connected, one of them changes to the circuit index of the other.

	//part 1
	fmt.Printf("The part1 answer is %d\n", answer)
}
