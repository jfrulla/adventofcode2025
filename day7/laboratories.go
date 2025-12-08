package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tachyonWorlds struct {
}

// MAIN
func main() {
	// Slurp up the input file
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// get our datastructures ready
	manifoldDiagram := [][]rune{}
	worldDiagram := [][]int{}
	splits := 0

	// For each line in the input file
	lines := strings.Split(string(inputFileContents), "\n")
	for lineNum, line := range lines {
		if line == "" {
			continue
		}
		manifoldDiagram = append(manifoldDiagram, []rune(line))
		worldDiagram = append(worldDiagram, make([]int, len(line)))
		if lineNum == 0 {
			continue
		}
		for charIndex, char := range line {
			if char == '.' {
				if manifoldDiagram[lineNum-1][charIndex] == 'S' {
					manifoldDiagram[lineNum][charIndex] = '|'
					worldDiagram[lineNum][charIndex] = 1
				}
				if manifoldDiagram[lineNum-1][charIndex] == '|' {
					manifoldDiagram[lineNum][charIndex] = '|'
					worldDiagram[lineNum][charIndex] += worldDiagram[lineNum-1][charIndex]
				}
			}
			if char == '^' && manifoldDiagram[lineNum-1][charIndex] == '|' {
				splits++
				manifoldDiagram[lineNum][charIndex-1] = '|'
				manifoldDiagram[lineNum][charIndex+1] = '|'
				worldDiagram[lineNum][charIndex-1] += worldDiagram[lineNum-1][charIndex]
				worldDiagram[lineNum][charIndex+1] += worldDiagram[lineNum-1][charIndex]
			}
		}
	}

	worlds := 0
	for _, world := range worldDiagram[len(worldDiagram)-1] {
		worlds += world
	}

	for _, line := range manifoldDiagram {
		fmt.Printf("%s\n", string(line))
	}
	for _, line := range worldDiagram {
		for _, world := range line {
			fmt.Printf("%s", strconv.Itoa(world))
		}
		fmt.Printf("\n")
	}

	//part 1
	fmt.Printf("The part1 answer is %d\n", splits)

	//part 2
	fmt.Printf("The part2 answer is %d\n", worlds)

}

// at every position, each beam has a number of possible worlds associated with it
//     S
//     1
//   1 ^ 1
//   1   1
// 1 ^ 2 ^ 1
// 1   2   1
//1^ 3 ^ 3 ^ 1, sum = 8

// OH THATS A BINOMIAL COEFFICIENT SEQUENCE!!
// but with a slight variation on splits vs non-splits
