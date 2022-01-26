// Map is a collection where each value is accessed via a key.
// Keys are an easy way to get data back out of your map.
// It's like having neatly labeled file folders instead of a messy pile.
package main

import "fmt"

func main() {
	// myMap is a map variable.
	// To declare a variable that holds a map, you type the map keyword, followed by square brackets [] containing the key type.
	// Then following the brackets, provide the value type.
	//var myMap map[string]float64

	// Just as with slices, declaring a map variable doesn't automatically create a map.
	// You need to call the "make" function.
	var ranks map[string]int
	ranks = make(map[string]int)

	// It's easier to just use a short variable declaration.
	//ranksShort := make(map[string]int)

	// Arrays and slices only let you use integer indexes.
	// You can choose almost any type to use for a map's keys.
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"]) // 3
	fmt.Println(ranks["gold"])   // 1

	// Another map with strings as keys and strings as values.
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"]) // Lithium
	fmt.Println(elements["H"])  // Hydrogen

	// Here's a map with integers as keys and booleans as values.
	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4]) // false
	fmt.Println(isPrime[7]) // true

	// Map literals are great if you know the keys and values you want your map to have in advance.
	pokemon := map[string]int{"pikachu": 1, "bulbasaur": 2, "geodude": 3}
	fmt.Println(pokemon["pikachu"]) // 1
	fmt.Println(pokemon["geodude"]) // 3
	flies := map[string]string{
		"winter": "midge",
		"spring": "leech",
		"summer": "mayfly",
		"fall":   "thai chub",
	}
	fmt.Println(flies["winter"]) // midge
	fmt.Println(flies["summer"]) // mayfly

	// If you leave a map literal empty, it starts empty and won't show an error.
	emptyMap := map[string]float64{}
	fmt.Println(emptyMap[""])
}
