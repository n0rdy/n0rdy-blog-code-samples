package main

import (
	"fmt"
	"time"
)

func main() {
	go printNTimes("Hello there.", 5)
	printNTimes("General Kenobi. You are a bold one.", 5)

	// to prevent the program from exiting before goroutine finishes
	time.Sleep(1 * time.Second)
}

func printNTimes(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(s)
	}
}
