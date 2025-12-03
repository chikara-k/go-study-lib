package main

import "fmt"

func goroutine1(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		ch <- sum
	}
	close(ch)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i)
	}
	/*
		go goroutine1(s, c)
		x := <-c
		fmt.Println(x)
		y := <-c
		fmt.Println(y)
	*/
}
