package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	password := 0
	dial := 50

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Open the input file
	fullInputPath := filepath.Join(cwd, "input")
	fileHandle, err := os.Open(fullInputPath)
	if err != nil {
		log.Fatal(err)
	}

	// For each line in the input file
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // skip empty lines
		}

		// break up the line into a direction and a number of clicks
		direction := string(line[0])
		number, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatal(err)
		}

		// Increment for each full rotation
		rotations := number / 100
		remainder := number % 100
		password += rotations

		// rotate the dial left or right and see if hit zero at any point
		if direction == "L" {
			if remainder >= dial && dial != 0 {
				password++
			}
			dial -= number
		} else if direction == "R" {
			if remainder >= (100-dial) && dial != 0 {
				password++
			}
			dial += number
		} else {
			log.Fatal("Could not interpret input correctly")
		}

		// correct the dial result
		dial %= 100
		if dial < 0 {
			dial += 100
		}
	}

	fmt.Printf("The password is %d\n", password)

}
