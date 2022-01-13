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

	// Slice operator - every slice is built on top of an underlying array.
	// Slice operator takes 2 indexes, the first is the starting point of the array and the second ends but does not include the index of the array.
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)
	// [a b c]

	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(slice2)
	// [b c d]

	// If you want a slice to include the last element of an underlying array, you specify a second index that's one beyond the end of the array.
	slice3 := underlyingArray[0:5]
	fmt.Println(slice3)
	// [a b c d e]

	// Slice operator has defaults for both the start and stop indexes.
	// If you omit the start index, a value of 0 (the first element of the array)
	slice4 := underlyingArray[:3]
	fmt.Println(slice4)
	// [a b c]

	// If you omit the stop index, the remaining underlying array is included in the resulting slice.
	slice5 := underlyingArray[1:]
	fmt.Println(slice5)
	// [b c d e]

	// Slices are just a view of the underlying array.
	// Think of slices as a microscope only being able to look at a portion of the contents of a slide.
	// Slices can overlap as well, slices can be a view into its own subset of the array's elements.
	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice6 := array3[0:3]
	slice7 := array3[2:5]
	fmt.Println(slice6, slice7)
	// [a b c] [c d e]

	// NOTE: Be careful about a slice.
	// A slice is just a view into the contents of an array.
	// If you change the underlying array, those changes will also be visible within the slice!
	ogArray := [5]string{"Chappelle", "Tupac", "Biggie", "Tiger", "Tony"}
	rappers := ogArray[1:3]
	fmt.Println(rappers)
	// [Tupac Biggie]
	// If you change the original underlying array ogArray to "X" at the second element, you change the view of the slice.
	ogArray[1] = "X"
	fmt.Println(rappers)
	// [X Biggie]

	// Since it is a problem that if you change the underlying array, the slice view changes.
	// It is recommended to create an array using the "make" function or with a slice literal instead.
}
