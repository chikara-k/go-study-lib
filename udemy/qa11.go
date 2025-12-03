package main

import "fmt"

type Employee struct {
	name     string
	age      int
	position string
}

func (e Employee) Introduce() {
	fmt.Printf("Hello, my name is %s, I am %d years old, and I work as a %s.\n", e.name, e.age, e.position)
}

func main() {
	e := Employee{name: "Alice", age: 28, position: "Software Engineer"}
	e.Introduce()
}
