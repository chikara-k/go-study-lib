package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	// forで回すようなときはclose(ch)でチャネルを閉じる
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}

	/*
		x := <-ch
		fmt.Println(x)
		fmt.Println(len(ch))

		ch <- 300
		fmt.Println(len(ch))
	*/
}
