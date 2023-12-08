package main

import "fmt"

func main() {
	ch := make(chan int)
	defer close(ch)

	go add(ch, 1, 2)

	result := <-ch
	fmt.Println(result)
}

func add(ch chan int, a int, b int) {
	ch <- a + b
}
