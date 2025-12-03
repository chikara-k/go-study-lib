package main

import "fmt"

// ...で可変長引数になる
func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)

	// スライスの中身を展開しながら入れれる
	foo(s...)
}
