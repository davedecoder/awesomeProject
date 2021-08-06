package main

import "fmt"

func ReverseInt(s []int) []int {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}

func ReverseString(s []string) []string {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}

func main() {
	a := []int{1,2,3}
	b := []string{"1","2","3"}
	fmt.Println(ReverseInt(a))
	fmt.Println(ReverseString(b))
	fmt.Println("exit")
}