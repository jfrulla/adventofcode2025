package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	password := 0
	dial := 50
	fileHandle, _ := os.Open("input")
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		number, _ := strconv.Atoi(strings.TrimSpace(scanner.Text())[1:])
		if string(strings.TrimSpace(scanner.Text())[0]) == "L" {
			dial -= number
		} else {
			dial += number
		}
		dial %= 100
		if dial < 0 {
			dial += 100
		}
		if dial == 0 {
			password++
		}
	}
	fmt.Printf("The password is %d\n", password)
}
