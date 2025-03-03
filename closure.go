package main

import "fmt"

// 　関数を返す関数
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

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
	fmt.Println(c1(3)) // 3.14を保持する

	c2 := circleArea(3)
	fmt.Println(c2(2))
	fmt.Println(c2(3)) // 3を保持する
}
