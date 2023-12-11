package main

import "fmt"

func main() {
	zeroCapacityChan := make(chan string)

	zeroCapacityChan <- "a"

	received := <-zeroCapacityChan
	fmt.Println(received)
}
