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
	fmt.Printf("Desk worker %s passed %s to the back office\n", dw.Name, item)
}

type WizardBackOfficeWorker struct {
	Name string
}

func (bow *WizardBackOfficeWorker) StartWorkingDay(backOfficeDeskChan chan string, shutdownChan chan struct{}) {
	wg := &sync.WaitGroup{}

	for {
		select {
		case item := <-backOfficeDeskChan:
			fmt.Printf("Wizard %s received %s\n", bow.Name, item)

			wg.Add(1)
			go func(item string) {
				defer wg.Done()

				fmt.Printf("Wizard %s casted a spell to process %s\n", bow.Name, item)
				bow.Process(item)
			}(item)
		case <-shutdownChan:
			fmt.Printf("the back office is closed - time to go home, %s\n", bow.Name)
			wg.Wait()
			return
		}
	}
}

func (bow *WizardBackOfficeWorker) Process(item string) {
	fmt.Printf("Wizard %s's spell started processing %s...\n", bow.Name, item)

	// to simulate long processing
	time.Sleep(10 * time.Second)

	fmt.Printf("Wizard %s's spell finished processing %s\n", bow.Name, item)
}

func main() {
	start := time.Now()

	deskChan := make(chan string)
	backOfficeDeskChan := make(chan string)
	phoneChan := make(chan string)
	deskShutdownChan := make(chan struct{})
	backOfficeDeskShutdownChan := make(chan struct{})
	postOfficeShutdownChan := make(chan struct{})

	// to simulate a long working day
	time.AfterFunc(5*time.Minute, func() {
		postOfficeShutdownChan <- struct{}{}
	})

	wg := &sync.WaitGroup{}

	bobWorker := DeskWorker{Name: "Bob", BackOfficeDeskChan: backOfficeDeskChan}
	wizardBackOfficeWorker := WizardBackOfficeWorker{Name: "Radagast"}

	wg.Add(1)
	go func() {
		bobWorker.StartWorkingDay(deskChan, phoneChan, deskShutdownChan)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		wizardBackOfficeWorker.StartWorkingDay(backOfficeDeskChan, backOfficeDeskShutdownChan)
		wg.Done()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		phoneChan <- "Has my package arrived?"
		time.Sleep(1 * time.Second)
		phoneChan <- "What about now?"
	}()

	queueChan := make(chan Customer)
	go func() {
		for {
			select {
			case <-postOfficeShutdownChan:
				fmt.Println("the post office is closed - time to go home")
				close(queueChan)
				return
			default:
				// to simulate a random customer arrival while the post office is open
				customer := generateCustomerWithRandomWait()
				fmt.Printf("%s enters the post office\n", customer.Name)
				queueChan <- customer
			}
		}
	}()

	for customer := range queueChan {
		deskChan <- customer.GiveAway()
	}

	deskShutdownChan <- struct{}{}
	backOfficeDeskShutdownChan <- struct{}{}

	wg.Wait()

	close(phoneChan)
	close(deskChan)
	close(backOfficeDeskChan)
	close(deskShutdownChan)
	close(backOfficeDeskShutdownChan)
	close(postOfficeShutdownChan)

	end := time.Now()
	fmt.Println("\n=====================================")
	fmt.Printf("Execution time in milliseconds: %v\n", end.Sub(start).Milliseconds())
}
