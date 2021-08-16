package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func (c Counter) WontIncrement() {
	c.total++
	c.lastUpdated = time.Now()
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("In doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("In doUpdateWrong:", c.String())
}

func main() {
	counter := Counter{}
	fmt.Println(counter.String())
	counter.Increment()
	fmt.Println(counter.String())
	counter.WontIncrement()
	fmt.Println(counter.String())
	doUpdateWrong(counter)
	fmt.Println(counter.String())
	doUpdateRight(&counter)
	fmt.Println(counter.String())
}