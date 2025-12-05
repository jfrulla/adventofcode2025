package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type PaperGrid struct {
	matrix [][]rune
	m_max  int
	n_max  int
}

// MAIN
func main() {
	// get our datastructures ready
	grid := PaperGrid{} //[][]rune{}

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
		grid.matrix = append(grid.matrix, []rune(line))
	}

	// this assumes a rectangular grid
	grid.m_max = len(grid.matrix[0]) - 1
	grid.n_max = len(grid.matrix) - 1

	totalAccessible := 0
	accessibleRolls := getRolls(&grid, 1)
	fmt.Printf("The part1 answer is %d\n", accessibleRolls)

	for accessibleRolls > 0 {
		accessibleRolls = getRolls(&grid, 2)
		totalAccessible += accessibleRolls
	}
	fmt.Printf("The part2 answer is %d\n", totalAccessible)
}

// getRolls takes a pointer to a slice of strings and the max number
// of adjacent rolls as input and returns the total number of
// accessible rolls as an int
func getRolls(grid *PaperGrid, part int) int {
	count := 0

	for n := 0; n <= (*grid).n_max; n++ {
		for m := 0; m <= (*grid).m_max; m++ {

			// if the one we're looking at isn't a paper roll
			// then we don't care about it.
			if (*grid).matrix[n][m] != '@' {
				continue
			}

			adjacent := 0

			//top
			if n > 0 {
				//left
				if m > 0 {
					if (*grid).matrix[n-1][m-1] == '@' {
						adjacent++
					}
				}
				//middle
				if (*grid).matrix[n-1][m] == '@' {
					adjacent++
				}
				//right
				if m < (*grid).m_max {
					if (*grid).matrix[n-1][m+1] == '@' {
						adjacent++
					}
				}
			}
			// left
			if m > 0 {
				if (*grid).matrix[n][m-1] == '@' {
					adjacent++
				}
			}
			//right
			if m < (*grid).m_max {
				if (*grid).matrix[n][m+1] == '@' {
					adjacent++
				}
			}
			//bottom
			if n < (*grid).n_max {
				//left
				if m > 0 {
					if (*grid).matrix[n+1][m-1] == '@' {
						adjacent++
					}
				}
				//middle
				if (*grid).matrix[n+1][m] == '@' {
					adjacent++
				}

				//right
				if m < (*grid).m_max {
					if (*grid).matrix[n+1][m+1] == '@' {
						adjacent++
					}
				}
			}
			if adjacent < 4 {
				if part == 2 {
					(*grid).matrix[n][m] = '.'
				}
				count++
			}
		}
	}
	return count
}

// func evaluate cell
//func evalCell()
