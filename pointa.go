package main

import "fmt"

// &でアドレスを指定。*でアドレスが示す実態を指定。
func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)

	/*
		fmt.Println(n)

		// メモリのアドレスが表示される
		fmt.Println(&n)

		// *でポイント型。&nでアドレスを指定
		var p *int = &n

		fmt.Println(p)

		fmt.Println(*p)
	*/
}
