package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./switch.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()

	data := make([]byte, 100)
	// errは上書きされる
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	// エラーしか返さない場合もある
	err = os.Chdir("test")
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}
