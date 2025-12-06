package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type PaperGrid struct {
	matrix   [][]rune // character representation of the whole grid
	adjacent [][]int  // keep track of roll adjacency
	m_max    int      // lets not run len() more than once
	n_max    int      // lets not run len() more than once
	count    int      // keep count of removed rolls
}

var directions [][]int // A global defining adjacency vectors

// MAIN
func main() {
	// get our datastructures ready
	grid := PaperGrid{}
	directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

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
		grid.adjacent = append(grid.adjacent, make([]int, len(line)))
	}

	// this assumes a rectangular grid
	grid.m_max = len(grid.matrix[0]) - 1
	grid.n_max = len(grid.matrix) - 1

	getRollsPart1(&grid)
	fmt.Printf("The part1 answer is %d\n", grid.count)

	getRollsPart2(&grid)
	fmt.Printf("The part2 answer is %d\n", grid.count)
}

// getRollsPart1 takes a pointer to a struct containing
// relevant information about the grid, calculates
// how many rolls can be removed with a simple iteration
// and returns the count
func getRollsPart1(grid *PaperGrid) {
	(*grid).count = 0

	for n := 0; n <= (*grid).n_max; n++ {
		for m := 0; m <= (*grid).m_max; m++ {

			// if the one we're looking at isn't a paper roll
			// then we don't care about it.
			if (*grid).matrix[n][m] != '@' {
				continue
			}

			// count the rolls in the adjacent cells
			(*grid).adjacent[n][m] = 0
			for _, vector := range directions {
				nn, nm := n+vector[0], m+vector[1]
				if nn >= 0 && nn <= (*grid).n_max && nm >= 0 && nm <= (*grid).m_max {
					if (*grid).matrix[nn][nm] == '@' {
						(*grid).adjacent[n][m]++
					}
				}
			}

			if (*grid).adjacent[n][m] < 4 {
				(*grid).count++
			}
		}
	}
}

// getRollsPart2 takes a pointer to a struct containing
// relevant information about the grid, calculates
// how many rolls can be removed via indefinite iterations,
// and returns the count as an int
func getRollsPart2(grid *PaperGrid) {
	(*grid).count = 0

	// Call a recursive evaluation a minimum of one time per roll
	for n := 0; n <= (*grid).n_max; n++ {
		for m := 0; m <= (*grid).m_max; m++ {
			if (*grid).adjacent[n][m] < 4 && (*grid).matrix[n][m] == '@' {
				evalCell(grid, n, m)
			}
		}
	}
}

// evalCell takes a pointer to the grid, alongside a position vector,
// checks to see if it is a roll that can be removed, does or doesn't
// remove the roll, then calls itself recursively on relevant neighbors
func evalCell(grid *PaperGrid, n int, m int) {
	// if not a roll, then just don't even worry about it
	if (*grid).matrix[n][m] != '@' {
		return
	}

	// if it's a roll, then lets see how many are adjacent
	(*grid).adjacent[n][m] = 0
	for _, vector := range directions {
		nn, nm := n+vector[0], m+vector[1]
		if nn >= 0 && nn <= (*grid).n_max && nm >= 0 && nm <= (*grid).m_max {
			if (*grid).matrix[nn][nm] == '@' {
				(*grid).adjacent[n][m]++
			}
		}
	}

	// if we remove a roll, decrement adjacency in all nearby cells
	// and recursively check any nearby rolls that may qualify for removal
	if (*grid).matrix[n][m] == '@' && (*grid).adjacent[n][m] < 4 {
		(*grid).matrix[n][m] = 'x'
		(*grid).count++
		for _, vector := range directions {
			nn, nm := n+vector[0], m+vector[1]
			if nn >= 0 && nn <= (*grid).n_max && nm >= 0 && nm <= (*grid).m_max {
				(*grid).adjacent[nn][nm] -= 1
				if (*grid).adjacent[nn][nm] < 4 && (*grid).matrix[nn][nm] == '@' {
					evalCell(grid, nn, nm)
				}
			}
		}
	}
}
