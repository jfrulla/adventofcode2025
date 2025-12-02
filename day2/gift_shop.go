package main

// checking each number individually is O(n)
// but I'll be evaluating thousands or maybe millions of numbers
// I wonder if there's a trick in the parsing itself

func main() {
	// get the numbers in the ranges
	// split on the comma to get several sets of ranges
	// for each range, split on the hyphen to get first and last
	// i=first
	// for i <= last
	//		make the number a string
	//		check the number
	//		if mirrored, add to sum
	//		i++

	// numbers with an odd number of digits are simply valid
	// numbers with an even number of digits might be invalid
	// convert all remaining numbers to strings
	// break it evenly into 2 substrings
	// see if substring1 == substring2
	// add up all the invalid IDs
}
