package main

import (
	"fmt"
	"sync"
)

func printMessage() {
	defer wg.Done()
	fmt.Println("Hello from Goroutine!")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go printMessage()
	}

	wg.Wait()
}
