package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}
	numbers = append(numbers, 40)

	members := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	members["Charlie"] = 35

	fmt.Println(numbers)
	fmt.Println(members)
}
