package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "test"
}

func main() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("Mac!")
	case "Windows!":
		fmt.Println("Windows!")
	default:
		fmt.Println("Default!", os)
	}

	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!")
	case "Windows!":
		fmt.Println("Windows!")
	default:
		fmt.Println("Default!", os)
	}
	// os変数は使えない

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Print("Night")
	}
}
