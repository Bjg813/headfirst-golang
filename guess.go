// guess challenges players to guess a random number
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// turn time into a number with time.Now().Unix()
	seconds := time.Now().Unix()
	// Set the random seed to a new number each time using seconds variable
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	// Create a new keyboard reader with standard input
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		// record keyboard input when user hits enter and a new line is created
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// trim off the \n new line
		input = strings.TrimSpace(input)
		// convert the string to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		// create logic to track if user guessed to high or too low
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("You nailed it!")
			// Once guessed correctly break out of the conditional using the keyword "break"
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
