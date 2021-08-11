package main

import "fmt"

type person struct {
	age int
	name string
}

func main() {
	p := person{}
	i := 2
	s := "hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	sl := make([]int, 8)
	modSlice(sl)
	fmt.Println(sl)
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string){
	m[2] = "Hello"
	m[3] = "World"
	delete(m, 1)
}

func modSlice(s []int){
	for k, v := range s{
		s[k] = v + 2
	}
	s = append(s, 10)
}