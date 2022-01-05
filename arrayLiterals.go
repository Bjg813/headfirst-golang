package main

import "fmt"

func main() {
	// Notes is an array initialized using array literals.
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	fmt.Println(notes)
	fmt.Printf("%#v\n", notes)

	// Notes array initialized using the short declaration.
	notesShort := [7]string{"do", "re", "fa", "so", "la", "ti"}
	fmt.Println(notesShort[0], notesShort[1], notesShort[2])

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes[0], primes[1], primes[2], primes[3], primes[4])
	fmt.Println(primes)

	// Text is an array with 3 elements that are using a new line per string spaced between commas.
	// If you display your array literals on a new line, it must have a comma at the end.
	text := [3]string{
		"This is a series of long string",
		"which would be awkward to place",
		"together on a single line",
	}
	fmt.Println(text[0], text[1], text[2])
}
