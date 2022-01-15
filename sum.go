// Sum package takes as many numbers as you want and adds them together using a variadic sum function.
package main

import "fmt"

// Sum is a variadic function that can take as many int arguments as you need to add together and returns an int sum.
func sum(numbers ...int) int {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(sum(7, 9))
	fmt.Println(sum(4, 1, 2))
}
