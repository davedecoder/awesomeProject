package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"strconv"
	"log"
)

func parser() (bool, error) {
	flag, err := strconv.ParseBool("true")
	return flag, err
}

func privateFunction() {
	fmt.Println("This is a \"private\" function")
}

func say(phrase string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(phrase)
	}
}

func divide(numerator float64, denominator float64) (float64, error) {

	if denominator == 0 {
		return 0, fmt.Errorf("Denominator cannot be zero bitch!")
	}
	return numerator / denominator, nil
}

func main() {
	var myFloat = 3.1415
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))
	fmt.Println(strings.Title("My first GO program"))
	fmt.Println(math.Floor(3.1415))
	fmt.Println("Hello, 世界")
	if true {
		fmt.Println("hola mundo")
	}
	for i := 1; i <= int(myFloat); i++ {
		fmt.Println(fmt.Sprintf("index: %d ##", i))
	}
	privateFunction()
	fmt.Println(parser())

	quotient, err := divide(5.6, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f\n", quotient) // => 2.80
	quotient, err = divide(5.6, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f\n", quotient) // => 2.80

	// pointers
	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt

	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)

	flag, err := parser()
	if err != nil {
		log.Fatal(err)
	}
	if flag {
		fmt.Println("saxxxes")
	}

	primes := make([]int, 2)
	primes[0] = 2
	primes[1] = 3
	primes = append(primes, 4)
	primes = append(primes, 5)
	fmt.Println(primes)

}
