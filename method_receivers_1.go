package main

import "fmt"

type Age int

func (a *Age) Increment() {
	*a = *a + 15
}

func main() {
	var a = Age(5)
	fmt.Println(a)
	a.Increment()
	fmt.Println(a)
}
