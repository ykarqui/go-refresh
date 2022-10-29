package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	name, lastName, city string
	age                  int
}

// greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, I'm " + p.name + " " + p.lastName + " and I'm " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// Init person using struct
	person1 := Person{name: "Sam", lastName: "Morty", city: "Vatican", age: 45}
	person2 := Person{"Sam", "Morty", "Vatican", 45}

	fmt.Println(person1)
	fmt.Println(person2.city)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
}
