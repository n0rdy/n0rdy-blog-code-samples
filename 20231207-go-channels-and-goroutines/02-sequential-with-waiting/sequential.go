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

func (w *Worker) Process(item string) {
	fmt.Printf("Worker %s received %s\n", w.Name, item)
	fmt.Printf("Worker %s started processing %s...\n", w.Name, item)

	// to simulate long processing
	time.Sleep(1 * time.Minute)

	fmt.Printf("Worker %s processed %s\n\n", w.Name, item)
}

func main() {
	bobWorker := Worker{Name: "Bob"}

	zlatan := Customer{Name: "Zlatan", Item: "football"}
	ben := Customer{Name: "Ben", Item: "box"}
	jenny := Customer{Name: "Jenny", Item: "watermelon"}
	eric := Customer{Name: "Eric", Item: "teddy bear"}
	lisa := Customer{Name: "Lisa", Item: "basketball"}

	queue := []Customer{lisa, eric, jenny, ben, zlatan}

	for _, customer := range queue {
		item := customer.GiveAway()
		bobWorker.Process(item)
	}
}
