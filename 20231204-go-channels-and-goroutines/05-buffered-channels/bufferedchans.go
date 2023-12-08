package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	defer close(ch)

	ch <- 1
	ch <- 2

	result1 := <-ch
	result2 := <-ch

	fmt.Println(result1)
	fmt.Println(result2)
}
