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
	// Basics
	fmt.Println("Hello, World!")
	fmt.Println(math.Pi)
	fmt.Println(add(33, 11))
	fmt.Println(swap("hello", "house"))
	var a, b = swap("hello", "world")
	c, d := swap("hello", "world")
	fmt.Println(a, b, c, d)

	// Flow
	sum := 0
	for i:=0; i < 10; i++ {
		sum +=1
	}
	fmt.Println(sum)

	if v := 2; v < 6 {
		fmt.Println("Yoshi")
	}

	// More Types: structs, slice, map
	type DataType struct {
		X int
		Y int
	}

	fmt.Println(DataType{1, 2})
	v := DataType{3, 45}
	v.X = 4
	fmt.Println(v.X, v.Y)

	var hollow [2]string
	hollow[0] = "Hollow"
	hollow[1] = "Knight"
	fmt.Println(hollow[0], hollow[1])
	fmt.Println(hollow)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[2:4])
	nums := []int{1, 2, 3}
	fmt.Println(append(nums, 4, 5))

	commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
  }
	commits["gri"] = 222
	fmt.Println(commits)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
}
