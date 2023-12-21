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

func (w *Worker) StartWorkingDay(deskChan chan string, phoneChan chan string, shutdownChan chan struct{}) {
	for {
		select {
		case item := <-deskChan:
			w.Process(item)
		case call := <-phoneChan:
			fmt.Printf("Worker %s received a call: %s\n", w.Name, call)
		case <-shutdownChan:
			fmt.Println("the desk is closed - time to go home")
			return
		}
	}
}

func (w *Worker) Process(item string) {
	fmt.Printf("Worker %s received %s\n", w.Name, item)
	fmt.Printf("Worker %s started processing %s...\n", w.Name, item)

	// to simulate long processing
	time.Sleep(1 * time.Second)

	fmt.Printf("Worker %s processed %s\n\n", w.Name, item)
}

func main() {
	start := time.Now()

	deskChan := make(chan string)
	phoneChan := make(chan string)
	shutdownChan := make(chan struct{})
	wg := &sync.WaitGroup{}

	bobWorker := Worker{Name: "Bob"}

	wg.Add(1)
	go func() {
		bobWorker.StartWorkingDay(deskChan, phoneChan, shutdownChan)
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

	shutdownChan <- struct{}{}

	close(deskChan)
	close(phoneChan)
	close(shutdownChan)

	wg.Wait()

	end := time.Now()
	fmt.Printf("Execution time in milliseconds: %v\n", end.Sub(start).Milliseconds())
}
