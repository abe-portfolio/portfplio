package main

import (
	"fmt"
)

// 返り値にfunc()を定義
// return時に関数の定義が可能
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// どういった時にclosureを使うか
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}

}

func main() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	c2 := circleArea(3)
	fmt.Println(c2(2))
}
