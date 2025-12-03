package main

import "fmt"

func numcheck(i int) string {
	if i%2 == 0 {
		return "偶数"
	} else {
		return "奇数"
	}
}

func main() {
	for i := 1; i <=100; i++ {
		fmt.Println(numcheck(i))
	}
}
