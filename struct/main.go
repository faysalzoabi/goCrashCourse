package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + strconv.Itoa(p.age)
}

// method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	// person1 := Person{firstName: "sama", lastName: "smith", city: "zurich", gender: "f", age: 24}
	person2 := Person{"sama", "smith", "zurich", "f", 24}

	// fmt.Println(person2)

	// fmt.Println(person2.firstName)
	// person2.age++

	person2.hasBirthday()
	person2.getMarried("Williams")
	fmt.Println(person2.greet())
	fmt.Println(person2.lastName)
}
