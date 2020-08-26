package main

import "fmt"

func main() {

	// var name = "Brad"
	var age int32 = 37
	var isCool = true
	isCool = false
	var size2 float32 = 2.4

	//shorthand
	// name := "Brad"
	// email := "brad@gmail.com"

	name, email := "faysal", "faysal@gmail.com"
	// size := 1.3
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size2)
}
