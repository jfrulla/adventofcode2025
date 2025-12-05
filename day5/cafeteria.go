package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IdRange struct {
	start int
	end   int
}

// MAIN
func main() {
	// Slurp up the input file
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// process & parse the input file
	ingredientIds, idRanges, err := processInput(inputFileContents)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// count the number of fresh ingredients we have for part1
	totalFresh := 0
	for _, id := range ingredientIds {
		for _, idRange := range idRanges {
			if id > idRange.start && id < idRange.end {
				totalFresh++
				break
			}
		}
	}
	fmt.Printf("The part1 answer is %d\n", totalFresh)

	// count the total number of possible valid IDs for part2
	idRanges = collapseRanges(idRanges)
	totalValid := 0
	for _, idRange := range idRanges {
		totalValid += idRange.end - idRange.start + 1
	}

	fmt.Printf("The part2 answer is %d\n", totalValid)
}

// procesInput takes the entire bytestring from the file and
// processes it into a slice of ingredientIDs and a slice of
// valid ID ranges
func processInput(fileInput []byte) ([]int, []IdRange, error) {
	// get our datastructures ready
	ingredientIds := []int{}
	idRanges := []IdRange{}

	// For each line in the input file
	processIngredientIds := false
	lines := strings.Split(string(fileInput), "\n")
	for _, line := range lines {

		// A blank line means that all following entries
		// are individual ingredientIds to check
		if line == "" {
			processIngredientIds = true
			continue
		}
		if processIngredientIds {
			id, err := strconv.Atoi(strings.TrimSpace(line))
			if err != nil {
				return nil, nil, fmt.Errorf("Could not convert %s to int: %w", line, err)
			}
			ingredientIds = append(ingredientIds, id)
		} else {
			parts := strings.Split(line, "-")
			firstNum, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, fmt.Errorf("Could not convert %s to int: %w", parts[0], err)
			}
			lastNum, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, fmt.Errorf("Could not convert %s to int: %w", parts[1], err)
			}
			idRanges = append(idRanges, IdRange{start: firstNum, end: lastNum})
		}
	}
	return ingredientIds, idRanges, nil
}

// collapseRanges takes the slice of ID ranges
// as input and returns a collapsed version of it
// where no ID ranges overlap.
func collapseRanges(idRanges []IdRange) []IdRange {
	newRanges := idRanges
	for _, idRange := range idRanges {
		tempRanges := []IdRange{}
		for _, tempRange := range newRanges {
			if idRange.start >= tempRange.start && idRange.start <= tempRange.end {
				idRange.start = tempRange.start
				idRange.end = max(idRange.end, tempRange.end)
			} else if idRange.end >= tempRange.start && idRange.end <= tempRange.end {
				idRange.end = tempRange.end
				idRange.start = min(idRange.start, tempRange.start)
			} else {
				tempRanges = append(tempRanges, tempRange)
			}
		}
		tempRanges = append(tempRanges, idRange)
		newRanges = tempRanges
	}
	return newRanges
}
