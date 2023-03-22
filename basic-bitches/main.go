package main

import "fmt"

type Number interface {
	int64 | float64
}

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

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// Type arguments can be omitted in calling code when the compiler can
	// infer the types you want to use. The compiler infers the types from
	// the types of function arguments passed on to it.
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums, with constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
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

// SumIntsOrFloats sums up the value of map m.
// Supports both int64 and float 64
// Generic function format func name [type parameters] (ordinary function parameters) return type {}
// int64 | float64 means the union of those types means that this constraint allows either types
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var accumulator V
	for _, v := range m {
		accumulator += v
	}
	return accumulator
}

// SumNumbers same as SumIntsOrFloats but uses the Number interface as a type constraint
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var accumulator V
	for _, v := range m {
		accumulator += v
	}
	return accumulator
}
