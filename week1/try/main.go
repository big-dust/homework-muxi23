package main

import "fmt"

func SumMap[K comparable, V int | float64](mp map[K]V) V {
	var sum V
	for _, v := range mp {
		sum += v
	}
	return sum
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Printf("Int: %v\n", SumMap(ints))
	fmt.Printf("Float: %v\n", SumMap(floats))
}
