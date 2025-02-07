package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	// int型を指定しいるのでintでキャストする
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
