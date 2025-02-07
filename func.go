package main

import "fmt"

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return result // resultを書かなくても良い（result int）になっていてresultが返るのがわかっているから
}

// func cal(price, item int) int {
// 	result := price * item
// 	return result
// }

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("Inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("Inner func", x)
	}(1)
}
