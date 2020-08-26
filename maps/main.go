package main

import "fmt"

func main() {
	//Define map
	// emails := make(map[string]string)

	// // Assign Kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	// declare map and key values
	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])

	// Delete from map
	delete(emails, "bob")
	fmt.Println(emails)
}
