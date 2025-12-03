package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

// *でポインタを指定することで特定メモリに格納されているstructを引数として取り込み、更新できる
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Area())
}
