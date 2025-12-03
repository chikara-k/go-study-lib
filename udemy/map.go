package main

import "fmt"

func main() {
	a := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(a)
	a["grape"] = 300
	fmt.Println(a)
	a["banana"] = 400
	fmt.Println(a["banana"])

	// 0
	fmt.Println(a["nothing"])

	// 0, false
	v, exist := a["nothing"]
	fmt.Println(v, exist)

	// まだメモリ上に展開されてないので errorになる
	/*
		var b map[string]int
		b["pc"] = 500
		fmt.Println(b)
	*/
}
