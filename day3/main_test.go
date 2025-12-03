package main

import "testing"

// main_test.go
func BenchmarkMainFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Call the main function or its logic here
		main()
	}
}
