package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

func main() {
	p := Person{name: "Taro", age: 30}
	p.Introduce()
}
