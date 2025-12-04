package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// MAIN
func main() {

	// get our datastructures ready
	grid := [][]rune{}

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
		grid = append(grid, []rune(line))
	}

	// this assumes a rectangular grid
	max_m := len(grid[0]) - 1
	max_n := len(grid) - 1

	totalAccessible := 0
	accessibleRolls := getRolls(&grid, max_m, max_n, 1)
	fmt.Printf("The part1 answer is %d\n", accessibleRolls)

	for accessibleRolls > 0 {
		accessibleRolls = getRolls(&grid, max_m, max_n, 2)
		totalAccessible += accessibleRolls
	}
	fmt.Printf("The part2 answer is %d\n", totalAccessible)
}

// getRolls takes a pointer to a slice of strings and the max number
// of adjacent rolls as input and returns the total number of
// accessible rolls as an int
func getRolls(grid *[][]rune, max_m int, max_n int, part int) int {
	count := 0

	for n := 0; n <= max_n; n++ {
		for m := 0; m <= max_m; m++ {

			// if the one we're looking at isn't a paper roll
			// then we don't care about it.
			if (*grid)[n][m] != '@' {
				continue
			}

			adjacent := 0

			//top
			if n > 0 {
				//left
				if m > 0 {
					if (*grid)[n-1][m-1] == '@' {
						adjacent++
					}
				}
				//middle
				if (*grid)[n-1][m] == '@' {
					adjacent++
				}
				//right
				if m < max_m {
					if (*grid)[n-1][m+1] == '@' {
						adjacent++
					}
				}
			}
			// left
			if m > 0 {
				if (*grid)[n][m-1] == '@' {
					adjacent++
				}
			}
			//right
			if m < max_m {
				if (*grid)[n][m+1] == '@' {
					adjacent++
				}
			}
			//bottom
			if n < max_n {
				//left
				if m > 0 {
					if (*grid)[n+1][m-1] == '@' {
						adjacent++
					}
				}
				//middle
				if (*grid)[n+1][m] == '@' {
					adjacent++
				}

				//right
				if m < max_m {
					if (*grid)[n+1][m+1] == '@' {
						adjacent++
					}
				}
			}
			if adjacent < 4 {
				if part == 2 {
					(*grid)[n][m] = '.'
				}
				count++
			}
		}
	}
	return count
}
