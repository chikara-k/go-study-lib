package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Hello World"
	fmt.Println(string(s[0]))

	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println(`Test
Test`)
}
