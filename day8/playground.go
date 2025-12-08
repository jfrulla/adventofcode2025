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
	circuits []int
}

// MAIN
func main() {
	// Slurp up the input file
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// get our datastructures ready
	junctionBoxCoords := [][]int{} // it'll be a list of {x, y, z} coordinates
	circuits := [][]int{}

	answer := 0

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
		circuits = append(circuits, []int{index})
	}
	edges := getDistances(junctionBoxCoords)
	connectBoxes(&edges, &circuits, 1000)

	fmt.Printf("Circuits: %v\n", circuits)

	threeLargest := make([]int, 3)

	for _, boxList := range circuits {
		if len(boxList) >= threeLargest[0] {
			threeLargest[2] = threeLargest[1]
			threeLargest[1] = threeLargest[0]
			threeLargest[0] = len(boxList)
		} else if len(boxList) >= threeLargest[1] {
			threeLargest[2] = threeLargest[1]
			threeLargest[1] = len(boxList)
		} else if len(boxList) >= threeLargest[2] {
			threeLargest[2] = len(boxList)
		}
	}
	answer = threeLargest[0] * threeLargest[1] * threeLargest[2]

	//part 1
	fmt.Printf("The part1 answer is %d\n", answer)
}

func getDistances(junctionBoxCoords [][]int) []Edge {
	edges := []Edge{}
	for box1Num, _ := range junctionBoxCoords {
		for box2Num := box1Num + 1; box2Num < len(junctionBoxCoords); box2Num++ {
			// https://en.wikipedia.org/wiki/Euclidean_distance
			boxDistance := math.Sqrt(math.Pow(float64(junctionBoxCoords[box2Num][0]-junctionBoxCoords[box1Num][0]), 2.0) + math.Pow(float64(junctionBoxCoords[box2Num][1]-junctionBoxCoords[box1Num][1]), 2.0) + math.Pow(float64(junctionBoxCoords[box2Num][2]-junctionBoxCoords[box1Num][2]), 2.0))
			edges = append(edges, Edge{distance: boxDistance, circuits: []int{box1Num, box2Num}})
		}
	}

	//sort the slice of edges based on distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	return edges
}

func connectBoxes(edges *[]Edge, circuits *[][]int, numConnections int) {
	for i := 0; i < numConnections; i++ {
		fmt.Printf("connecting %v\n", (*edges)[i].circuits)
		connect((*edges)[i].circuits, edges, circuits)
	}
}

func connect(boxes []int, edges *[]Edge, circuits *[][]int) {
	// first check to see if the boxes are in the same circuit
	// if so, do nothing
	if checkCircuitBoxes(boxes, circuits) {
		fmt.Printf("Already same circuit\n")
		return
	} else {
		mergeCircuits(boxes, edges, circuits)
	}
}

// returns true if same circuit already
// returns false if different circuits
func checkCircuitBoxes(boxes []int, circuits *[][]int) bool {
	box1 := boxes[0]
	box2 := boxes[1]
	for _, circuit := range *circuits {
		box1Found := false
		box2Found := false
		for _, box := range circuit {
			if box == box1 {
				box1Found = true
			}
			if box == box2 {
				box2Found = true
			}
		}
		if box1Found && box2Found {
			return true
		}
	}
	return false
}

func mergeCircuits(boxes []int, edges *[]Edge, circuits *[][]int) {
	box1 := boxes[0]
	box2 := boxes[1]
	box1Circuit := 0
	box2Circuit := 1
	// merge the circuit containing box2 into the circuit containing box1
	//loop through each circuit looking for box1 and box2
	for circuitIndex, circuit := range *circuits {
		// for each box in the circuit
		for _, box := range circuit {
			if box == box1 {
				box1Circuit = circuitIndex
			}
			if box == box2 {
				box2Circuit = circuitIndex
			}
		}
	}
	(*circuits)[box1Circuit] = append((*circuits)[box1Circuit], (*circuits)[box2Circuit]...)
	(*circuits)[box2Circuit] = []int{}
}
