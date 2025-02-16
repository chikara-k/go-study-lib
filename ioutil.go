package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("ioutil.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
		log.Fatalln(err)
	}
}
