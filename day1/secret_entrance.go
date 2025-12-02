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

		fmt.Printf("Interpreting %s\n", line)

		// break up the line into a direction and a number of clicks
		direction := string(line[0])
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		// rotate the dial left or right
		if direction == "L" {
			dial -= number
		} else if direction == "R" {
			dial += number
		} else {
			log.Fatal("Could not interpret input correctly")
		}

		// get the dial into the positive via increments of 100
		for dial < 0 {
			dial += 100
		}

		// modulus to get below 100
		dial %= 100

		// determine if we've hit zero and increment if we have
		if dial == 0 {
			password++
		}

		fmt.Printf("The dial is now at %d\n", dial)
	}

	fmt.Printf("The password is %d\n", password)

}
