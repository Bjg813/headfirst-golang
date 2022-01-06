package main

import "fmt"

func main() {
	// Use a integer variable to find the value at a particular index of an array.
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	index := 1
	fmt.Println(index, notes[index])
	index = 2
	fmt.Println(index, notes[index])

	// Use a for loop to iterate through an array.
	for i := 0; i <= 6; i++ {
		fmt.Println(i, notes[i])
	}

	// Use "len" to check the length of a function.
	fmt.Println("length of notes array is:", len(notes))

	// Use "len" to process an entire array to determine which indexes are safe to access.
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	// Use the for...range loop to safely access an array without risking accidentally accessing an invalid array length.
	// The for...range loop is safer and easier to read.
	// We provide the for...range loop with a variable for the index and a variable for the element.
	for index, note := range notes {
		fmt.Println("using for...range loop", index, note)
	}

	// Sharks array holds four string shark names.
	sharks := [4]string{"whale", "tiger", "blue", "mako"}

	// Use for...range loop to safely access and return all shark names in the sharks array.
	for index, shark := range sharks {
		fmt.Println(index, shark)
	}
}
