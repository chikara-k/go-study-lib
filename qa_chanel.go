package main

import (
	"fmt"
)

// 送信 (ch <- value) → チャネルにデータを送る
// 受信 (value := <-ch) → チャネルからデータを受け取る

func runTask(i int, ch chan string) {
	ch <- fmt.Sprintf("Task %i completed\n", i)
}

func main() {
	ch := make(chan string, 3)

	for i := 1; i <= 3; i++ {
		runTask(i, ch)
	}

	for i := 1; i <= 3; i++ {
		msg := <-ch
		fmt.Println(msg)
	}

	ch.Close()
}
