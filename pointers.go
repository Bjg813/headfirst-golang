package main

import (
	"fmt"
)

func main() {
	myInt := 4
	myIntPointer := &myInt
	*myIntPointer = 8
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
}
