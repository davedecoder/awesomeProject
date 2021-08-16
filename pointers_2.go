package main

import (
	"fmt"
	"reflect"
)

type person_1 struct {
	FirstName string
	MiddleName *string
	LastName string
}

func stringToPointer(input string) *string {
	return &input
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func update(px *int){
	*px = 20
}

func main() {
	p := person_1{
		"dave",
		stringToPointer("ivan"),
		"ruiz",
	}
	fmt.Println(p)
	fmt.Println(*p.MiddleName)
	fmt.Println(&p)

	x := 11
	f := &x
	failedUpdate(f)
	fmt.Println(f)
	fmt.Println(*f)
	update(f)
	fmt.Println(f)
	fmt.Println(x)

	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(x))
}
