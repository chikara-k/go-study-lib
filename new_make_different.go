package main

import "fmt"

func main() {
	// newでメモリを確保して、値には0が入る
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	// メモリを確保してないのでnilになって、nilに++したらpanicになる
	/*
		var p2 *int
		fmt.Println(p2)
		*p2++
		fmt.Println(p2)
	*/

	// ポイントを返すか返さないかによって、初期化にnewを使うかmakeを使うかことなる
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p3 *int = new(int)
	fmt.Printf("%T\n", p3)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
