package main

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
