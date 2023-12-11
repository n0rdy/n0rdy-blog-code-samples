package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printNTimesAsync("Hello there.", 5, wg)
		printNTimes("General Kenobi. You are a bold one.", 5)
	}

	// to prevent the program from exiting before goroutine finishes
	wg.Wait()
}

func printNTimes(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(s)
	}
}

func printNTimesAsync(s string, n int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		fmt.Println(s)
	}
	wg.Done()
}
