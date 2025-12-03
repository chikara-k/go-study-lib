package main

import (
	"fmt"
	"os"
)

func hello() {
	defer fmt.Println("world")

	fmt.Println("Hello")
}

func main() {
	/*
		hello()
		defer fmt.Println("after")

		fmt.Println("before")
	*/

	/*
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("finish")
	*/

	file, _ := os.Open("./switch.go")
	defer file.Close()
	data := make([]byte, 500)
	file.Read(data)
	fmt.Println(string(data))
}
