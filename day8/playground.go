package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	distance float64
	boxes    []int
}

// MAIN
func main() {
	// Slurp up the input
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// get our datastructures ready
	junctionBoxCoords := [][]int{} // it'll be a list of {x, y, z} coordinates
	circuits := [][]int{}

	// For each line in the input file
	lines := strings.Split(string(inputFileContents), "\n")
	for index, line := range lines {
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
		circuits = append(circuits, []int{index}) // one circuit per box to start
	}

	// build the graph
	edges := getDistances(junctionBoxCoords)

	// Note: from this point forward, we mostly only
	// reference boxes by their index in junctionBoxCoords

	// PART 1
	// make the 1000 closest connections
	for i := 0; i < 1000; i++ {
		connect(edges[i].boxes, &circuits)
	}

	// sort slice of circuits by descending length
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	// multiply the length of the 3 longest circuits
	answer := len(circuits[0]) * len(circuits[1]) * len(circuits[2])
	fmt.Printf("The part1 answer is %d\n", answer)

	// PART 2
	lastConnectedPair := []int{}

	// make connections until we only have 1 circuit
	for i := 0; len(circuits) > 1; i++ {
		connect(edges[i].boxes, &circuits)
		lastConnectedPair = edges[i].boxes
	}

	// multiply the x-values of the last 2 boxes that were connected together
	answer = junctionBoxCoords[lastConnectedPair[0]][0] * junctionBoxCoords[lastConnectedPair[1]][0]
	fmt.Printf("The part2 answer is %d\n", answer)
}

// getDistances will find the distance between every pair of boxes,
// create the edge object and then return the slice of all edges
// sorted by the distance between the pair
func getDistances(junctionBoxCoords [][]int) []Edge {
	edges := []Edge{}
	for box1Num, _ := range junctionBoxCoords {
		for box2Num := box1Num + 1; box2Num < len(junctionBoxCoords); box2Num++ {
			// https://en.wikipedia.org/wiki/Euclidean_distance
			boxDistance := math.Sqrt(math.Pow(float64(junctionBoxCoords[box2Num][0]-junctionBoxCoords[box1Num][0]), 2.0) + math.Pow(float64(junctionBoxCoords[box2Num][1]-junctionBoxCoords[box1Num][1]), 2.0) + math.Pow(float64(junctionBoxCoords[box2Num][2]-junctionBoxCoords[box1Num][2]), 2.0))
			edges = append(edges, Edge{distance: boxDistance, boxes: []int{box1Num, box2Num}})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})
	return edges
}

// connect takes a pair of box indices and a ptr to the slice of all circuits.
// If the pair of boxes is in the same circuit, does nothing.
// If they are two separate circuits, the circuits will be merged together.
func connect(boxes []int, circuits *[][]int) {
	if checkCircuitBoxes(boxes, circuits) {
		return
	} else {
		mergeCircuits(boxes, circuits)
	}

}

// checkCircuitBoxes returns true if the box indices are in the same circuit already.
// returns false if they are in different circuits.
func checkCircuitBoxes(boxes []int, circuits *[][]int) bool {
	for _, circuit := range *circuits {
		box1Found := false
		box2Found := false
		for _, box := range circuit {
			if box == boxes[0] {
				box1Found = true
			}
			if box == boxes[1] {
				box2Found = true
			}
		}
		if box1Found && box2Found {
			return true
		}
	}
	return false
}

// mergeCircuits will find the two circuits associated with the
// given pair of boxes and will then collapse one circuit into another
func mergeCircuits(boxes []int, circuits *[][]int) {
	box1Circuit, box2Circuit := 0, 0
	for circuitIndex, circuit := range *circuits {
		for _, box := range circuit {
			if box == boxes[0] {
				box1Circuit = circuitIndex
			}
			if box == boxes[1] {
				box2Circuit = circuitIndex
			}
		}
	}
	(*circuits)[box1Circuit] = append((*circuits)[box1Circuit], (*circuits)[box2Circuit]...)
	if box2Circuit < len(*circuits)-1 {
		*circuits = append((*circuits)[:box2Circuit], (*circuits)[box2Circuit+1:]...)
	} else {
		*circuits = (*circuits)[:box2Circuit]
	}
}
