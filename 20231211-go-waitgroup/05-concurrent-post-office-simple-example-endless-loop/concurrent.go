package main

import (
	"fmt"
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

func (w *Worker) StartWorkingDay(deskChan chan string) {
	for {
		item := <-deskChan
		w.Process(item)
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

	bobWorker := Worker{Name: "Bob"}
	go bobWorker.StartWorkingDay(deskChan)

	zlatan := Customer{Name: "Zlatan", Item: "football"}
	ben := Customer{Name: "Ben", Item: "box"}
	jenny := Customer{Name: "Jenny", Item: "watermelon"}
	eric := Customer{Name: "Eric", Item: "teddy bear"}
	lisa := Customer{Name: "Lisa", Item: "basketball"}

	queue := []Customer{lisa, eric, jenny, ben, zlatan}

	for _, customer := range queue {
		deskChan <- customer.GiveAway()
	}

	close(deskChan)
}
