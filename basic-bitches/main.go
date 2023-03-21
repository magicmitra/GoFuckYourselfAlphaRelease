package main

import "fmt"

func main() {
	// Initialize a map for integer vals
	ints := map[string]int64{
		"first":  69,
		"second": 77,
	}

	// Initialize a map for float vals
	floats := map[string]float64{
		"first":  69.99,
		"second": 77.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))
}

// SumInts adds together the values of m
func SumInts(m map[string]int64) int64 {
	var accumulator int64
	for _, v := range m {
		accumulator += v
	}
	return accumulator
}

// SumFloats adds together the values of m
func SumFloats(m map[string]float64) float64 {
	var accumulator float64
	for _, v := range m {
		accumulator += v
	}
	return accumulator
}
