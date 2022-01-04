// pass_fail reports whether a grade is passing or failing
// file block
package main

// package block

import (
	"fmt"
	"headfirst_golang/go/src/keyboard"
	"log"
)

func main() {
	// function block
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		// conditional block
		log.Fatal(err)
	}

	// declare "status" string variable in function block scope before using in conditional block
	var status string
	if grade > 60 {
		// conditional block
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
