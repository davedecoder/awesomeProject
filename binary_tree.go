package main

import "fmt"

type IntTree struct {
	val int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	}
	if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	if it == nil {
		return false
	}
	if val < it.val {
		return it.left.Contains(val)
	}
	if val > it.val {
		return it.right.Contains(val)
	}
	return true
	//switch {
	//case it == nil:
	//	return false
	//case val < it.
	//}
}

func main() {
	it := new(IntTree)
	it = it.Insert(5)
	it = it.Insert(5)
	it = it.Insert(6)
	it = it.Insert(7)
	it = it.Insert(2)
	it = it.Insert(1)
	it = it.Insert(10)
	it = it.Insert(0)
	fmt.Println(it.Contains(5))
	fmt.Println(it.Contains(1))
	fmt.Println(it.Contains(12))
	fmt.Println(*it.right)
}
