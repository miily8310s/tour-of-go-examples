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
}
