// Declare a slice use an empty pair of square brackets followed by type of elements the slice will hold.
//var mySlice []string.
package main

import "fmt"

func main() {
	// Unlike array variables, declaring a slice variable doesn't automatically create a slice.
	// You must call the built-in make function.
	var notes []string
	notes = make([]string, 7)

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "me"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	// You don't have to declare a slice in two seperate steps though, you can use a short variable.
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	notesSeven := make([]string, 7)
	primesFive := make([]int, 5)
	fmt.Println(len(notesSeven))
	fmt.Println(len(primesFive))

	// Both for and for...range loops work just the same with slices as they do with arrays.
	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(i, letters[i])
	}

	for _, letter := range letters {
		fmt.Println(letter)
	}

	// Slice literals are used the same way as array literals.
	notesLiteral := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notesLiteral[3], notesLiteral[6], notesLiteral[0])
	primesLiteral := []int{
		2,
		3,
		5,
	}
	fmt.Println(primesLiteral[0], primesLiteral[1], primesLiteral[2])
}
