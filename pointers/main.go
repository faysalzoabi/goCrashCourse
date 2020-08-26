package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)

	// Use * to read val from address

	fmt.Println(*&a)

	// change val with pointer
	*b = 10
	fmt.Println(a)

}
