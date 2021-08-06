package main

import (
	"fmt"
	"sort"
)

func makeMul(x int) func(y int) int{
	return func(y int) int {
		return x * y
	}
}

func main() {
	type Person struct {
		FirstName string
		LastName string
		Age int
	}

	people := []Person{
		{"Dave", "Marmotov", 32},
		{"Tere", "Lobatovsky", 31},
		{"Jochi", "Ruiz", 22},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	a := makeMul(2)
	b := makeMul(3)

	for i:=0; i<4; i++ {
		fmt.Println(a(i), b(i))
	}
}
