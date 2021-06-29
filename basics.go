package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Pi)
	fmt.Println(add(33, 11))
	fmt.Println(swap("hello", "house"))
	var a, b = swap("hello", "world")
	c, d := swap("hello", "world")
	fmt.Println(a, b, c, d)
}
