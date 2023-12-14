package main

import "fmt"

func main() {
	phoneChan := make(chan string)

	// listen to the phone calls:
	go func() {
		for call := range phoneChan {
			fmt.Println("Got a call: " + call)
		}
	}()

	// somebody calls the phone:
	phoneChan <- "Yo! What's up?"
	close(phoneChan)
}
