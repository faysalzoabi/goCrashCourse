package main

import "fmt"

func main() {
	//Arrays
	// var fruitArray [2]string

	// //Assign Values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// Declare and assign
	// fruitArray := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArray)

	fruitSlice := []string{"apple", "orange", "grape", "cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
	fmt.Println(fruitSlice[1:3])
}
