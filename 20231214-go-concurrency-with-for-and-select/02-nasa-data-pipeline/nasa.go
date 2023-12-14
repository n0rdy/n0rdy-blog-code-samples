package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type MarsRover struct{}

func (mr *MarsRover) GatherMetrics() int {
	fmt.Println("Mars rover: gathering metrics...")
	return rand.Int()
}

type Iss struct{}

func (i *Iss) Enrich(metrics int) string {
	fmt.Printf("ISS: enriching metrics [%d]...\n", metrics)
	return "ISS" + strconv.Itoa(metrics)
}

type NasaDataCenter struct{}

func (ndc *NasaDataCenter) Process(data string) {
	fmt.Printf("Nasa data center: processing data [%s]...\n", data)
}

func main() {
	marsRover := MarsRover{}
	iss := Iss{}
	nasaDataCenter := NasaDataCenter{}

	issChan := make(chan int)
	nasaChan := make(chan string)

	go func() {
		for metrics := range issChan {
			nasaChan <- iss.Enrich(metrics)
		}
	}()

	go func() {
		for data := range nasaChan {
			nasaDataCenter.Process(data)
		}
	}()

	for {
		issChan <- marsRover.GatherMetrics()
	}
}
