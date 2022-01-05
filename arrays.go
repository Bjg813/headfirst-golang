package main

import (
	"fmt"
	"time"
)

func main() {
	// Notes array of seven strings.
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])
	// When an array is created, all the values it contains are initialized to the zero value for the type the array holds.
	// If there are no values in a string array element than it is an empty string "".
	fmt.Println(notes[3])

	// Primes array of 5 integers.
	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(primes[1])
	// When an array is created, all the values it contains are initialized to the zero value for the type the array holds.
	// If there are no values in an integer array element than it is zero.
	fmt.Println(primes[2])

	// Dates array of 3 time values
	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	fmt.Println(dates[1])

	// Counters array of 3 counters
	// Counter is showing since int arrays initialize at zero, a counter can increment the element value up 1.
	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])
}
