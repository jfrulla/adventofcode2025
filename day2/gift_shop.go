package main

// checking each number individually is O(n)
// but I'll be evaluating thousands or maybe millions of numbers
// I wonder if there's a trick in the parsing itself

func main() {
	// get the numbers in the ranges
	// split on the comma to get several sets of ranges
	// for each range, split on the hyphen to get first and last
	// k = last - first
	// i=0
	// for i <= k
	//		search_nums.append(search_nums,first+i)
	//		i++
	// number with an odd number of digits are simply valid
	// convert all remaining numbers to strings
	// break it evenly into 2 substrings
	// see if substring1 == substring2
	// add up all the invalid IDs
}
