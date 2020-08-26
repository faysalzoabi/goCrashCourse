package main

import "fmt"

func main() {

	ids := []int{33, 45, 345, 23, 67, 32}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	for _, id := range ids {
		fmt.Printf("%d - ID:\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	// range with map
	emails := map[string]string{"fay": "syria", "idris": "swiss"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println(k)
	}
}
