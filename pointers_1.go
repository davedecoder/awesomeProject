package main

import "fmt"

func morePointers() {
	x := "hello"
	pointerToX := &x
	fmt.Println(pointerToX)

	y:= 10
	pointerToY := &y
	fmt.Println(pointerToY)
	fmt.Println(*pointerToY)
	z := 5 + *pointerToY
	fmt.Println(z)
}

func morePointers2() {
	var x = new(int)
	fmt.Println(x == nil) // prints false
	fmt.Println(*x)       // prints 0
}

func morePointers3() {
	type Foo struct {}
	var pointerToX *Foo
	fmt.Println(pointerToX)
	y := &Foo{}
	fmt.Println(&y)
}

func main() {
	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println(pointerX, pointerY, &pointerZ)
	morePointers()
	morePointers2()
	morePointers3()
}
