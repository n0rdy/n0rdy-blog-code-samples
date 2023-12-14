package main

import (
	"fmt"
	"sync"
	"time"
)

type Customer struct {
	Name string
	Item string
}

func (c *Customer) GiveAway() string {
	item := c.Item
	fmt.Printf("%s gives away %s\n", c.Name, item)
	c.Item = ""
	return item
}

type Worker struct {
	Name string
}

func (w *Worker) StartWorkingDay(deskChan chan string, phoneChan chan string) {
	keepRunning := true
	for keepRunning {
		select {
		case item, ok := <-deskChan:
			if ok {
				w.Process(item)
			} else {
				keepRunning = false
			}
		case call, ok := <-phoneChan:
			if ok {
				fmt.Printf("Worker %s received a call: %s\n", w.Name, call)
			} else {
				keepRunning = false
			}
		}
	}

	fmt.Println("the desk is closed - time to go home")
}

func (w *Worker) Process(item string) {
	fmt.Printf("Worker %s received %s\n", w.Name, item)
	fmt.Printf("Worker %s started processing %s...\n", w.Name, item)

	// to simulate long processing
	time.Sleep(1 * time.Second)

	fmt.Printf("Worker %s processed %s\n\n", w.Name, item)
}

func main() {
	deskChan := make(chan string)
	phoneChan := make(chan string)
	wg := &sync.WaitGroup{}

	bobWorker := Worker{Name: "Bob"}

	wg.Add(1)
	go func() {
		bobWorker.StartWorkingDay(deskChan, phoneChan)
		wg.Done()
	}()

	zlatan := Customer{Name: "Zlatan", Item: "football"}
	ben := Customer{Name: "Ben", Item: "box"}
	jenny := Customer{Name: "Jenny", Item: "watermelon"}
	eric := Customer{Name: "Eric", Item: "teddy bear"}
	lisa := Customer{Name: "Lisa", Item: "basketball"}

	queue := []Customer{lisa, eric, jenny, ben, zlatan}

	go func() {
		phoneChan <- "Has my package arrived?"
		time.Sleep(1 * time.Second)
		phoneChan <- "What about now?"
	}()

	for _, customer := range queue {
		deskChan <- customer.GiveAway()
	}

	close(deskChan)
	close(phoneChan)

	wg.Wait()
}
