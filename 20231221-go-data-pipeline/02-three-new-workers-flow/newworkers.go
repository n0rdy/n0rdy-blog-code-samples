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

type DeskWorker struct {
	Name               string
	BackOfficeDeskChan chan string
}

func (dw *DeskWorker) StartWorkingDay(deskChan chan string, phoneChan chan string, shutdownChan chan struct{}) {
	for {
		select {
		case item := <-deskChan:
			dw.Process(item)
		case call := <-phoneChan:
			fmt.Printf("Desk worker %s received a call: %s\n", dw.Name, call)
		case <-shutdownChan:
			fmt.Println("the desk is closed - time to go home")
			return
		}
	}
}

func (dw *DeskWorker) Process(item string) {
	fmt.Printf("Desk worker %s received %s\n", dw.Name, item)
	fmt.Printf("Desk worker %s started checking ID of the customer with the %s item...\n", dw.Name, item)

	// to simulate long processing
	time.Sleep(1 * time.Second)

	fmt.Printf("Desk worker %s finished checking ID of the customer with the %s item\n", dw.Name, item)

	fmt.Printf("Desk worker %s started passing %s to the back office...\n", dw.Name, item)
	dw.BackOfficeDeskChan <- item
	fmt.Printf("Desk worker %s passed %s to the back office\n\n", dw.Name, item)
}

type BackOfficeWorker struct {
	Name string
}

func (bow *BackOfficeWorker) StartWorkingDay(backOfficeDeskChan chan string, shutdownChan chan struct{}) {
	for {
		select {
		case item := <-backOfficeDeskChan:
			bow.Process(item)
		case <-shutdownChan:
			fmt.Printf("the back office is closed - time to go home, %s\n", bow.Name)
			return
		}
	}
}

func (bow *BackOfficeWorker) Process(item string) {
	fmt.Printf("Back office worker %s received %s\n", bow.Name, item)
	fmt.Printf("Back office worker %s started processing %s...\n", bow.Name, item)

	// to simulate long processing
	time.Sleep(10 * time.Second)

	fmt.Printf("Back office worker %s finished processing %s\n", bow.Name, item)
}

func main() {
	start := time.Now()

	deskChan := make(chan string)
	backOfficeDeskChan := make(chan string)
	phoneChan := make(chan string)
	deskShutdownChan := make(chan struct{})
	backOfficeDeskShutdownChan := make(chan struct{})
	wg := &sync.WaitGroup{}

	bobWorker := DeskWorker{Name: "Bob", BackOfficeDeskChan: backOfficeDeskChan}
	odaWorker := BackOfficeWorker{Name: "Oda"}
	robertWorker := BackOfficeWorker{Name: "Robert"}
	marthaWorker := BackOfficeWorker{Name: "Martha"}

	wg.Add(1)
	go func() {
		bobWorker.StartWorkingDay(deskChan, phoneChan, deskShutdownChan)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		odaWorker.StartWorkingDay(backOfficeDeskChan, backOfficeDeskShutdownChan)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		robertWorker.StartWorkingDay(backOfficeDeskChan, backOfficeDeskShutdownChan)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		marthaWorker.StartWorkingDay(backOfficeDeskChan, backOfficeDeskShutdownChan)
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

	deskShutdownChan <- struct{}{}
	// 3 stands for the number of back office workers
	for i := 0; i < 3; i++ {
		backOfficeDeskShutdownChan <- struct{}{}
	}

	close(phoneChan)
	close(deskChan)
	close(backOfficeDeskChan)
	close(deskShutdownChan)

	wg.Wait()

	end := time.Now()
	fmt.Println("\n=====================================")
	fmt.Printf("Execution time in milliseconds: %v\n", end.Sub(start).Milliseconds())
}
