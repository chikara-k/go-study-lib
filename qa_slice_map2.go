package main

import "fmt"

func main() {
	numbers := []int{5, 10, 15, 20, 25}
	evens := []int{}
	for _, number := range numbers {
		fmt.Println(number * 2)
		if number%2 == 0 {
			evens = append(evens, number)
		}
	}
	fmt.Println(evens)

	scores := map[string]int{
		"Alice":   85,
		"Bob":     90,
		"Charlie": 78,
	}
	scores["Bob"] = 95
	fmt.Println(scores["Bob"])

	delete(scores, "Charlie")
	if _, exist := scores["Charlie"]; !exist {
		fmt.Println("Charlie removed")
	}

	if _, exist := scores["David"]; !exist {
		fmt.Println("David not found")
	}
}
