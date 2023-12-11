package main

import "fmt"

func main() {
	oneCapacityChan := make(chan string, 1)

	oneCapacityChan <- "a"

	received := <-oneCapacityChan
	fmt.Println(received)
}
