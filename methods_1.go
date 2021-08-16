package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

func (p Person) String()string {
	return fmt.Sprintf("Person: %s %s, Age: %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	p := Person{
		"Dave",
		"Ruiz",
		32,
	}
	fmt.Println(p.String())
}
