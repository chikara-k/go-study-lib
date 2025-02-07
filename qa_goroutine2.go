package main

import (
	"fmt"
	"sync"
	"time"
)

// * と &はセットで使うことが一般的
// &（アドレス演算子） → 変数の メモリアドレス を取得する
// *（ポインタ演算子） → ポインタが指している値を取得・変更する

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg = sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed!")
}
