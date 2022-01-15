// Variadic function is one that can be called with a varying number of arguments by using an ellipsis(...).
/// To make a function variadic us an ellipsis(...) before the type of the last (or only) function paramenter in the function declaration.
package main

import (
	"fmt"
	"math"
)

// severalInts is a variadic function and works just fine with any number of arguments.
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

// severalStrings is a variadic function that can handle any number of arguments.
func severalStrings(strings ...string) {
	fmt.Println(strings)
}

// mix function can take one or more nonvariadic arguments as well as variadic arguments at the same time.
// Although a function caller can omit variadic arguments (resulting in an empty slice), nonvariadic arguments are always required; it's a compile error to omit those.
// Only the last parameter in a function definition can be variadic; you can't place it in front of required parameters.
func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	severalInts(1, 2, 3)
	// [1 2 3]

	severalStrings("how", "do", "you", "do")
	// [how do you do]

	// If there are no arguments, an empty slice is received and doesn't return an error.
	severalStrings()
	// []

	mix(1, true, "howdy", "ho")
	// 1 true [howdy ho]

	fmt.Println(maximum(71.8, 56.2, 89.5))
	// 89.5

	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
	//98.5

	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	// [3.2 50]

	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
	// [4.1 -5.2]
}
