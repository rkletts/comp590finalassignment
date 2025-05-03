package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	inCh := make(chan int, 5)   // stage 1 -> stage 2
	outCh := make(chan int, 5)  // stage 2 -> stage 3

	var wgProducers sync.WaitGroup
	var wgProcessors sync.WaitGroup

	// Stage 1: Producers
	wgProducers.Add(2)
	go producer("odd", 1, 29, 2, inCh, &wgProducers)
	go producer("even", 2, 30, 2, inCh, &wgProducers)

	// Stage 2: Processors
	wgProcessors.Add(2)
	go processor("proc1", inCh, outCh, &wgProcessors)
	go processor("proc2", inCh, outCh, &wgProcessors)

	// Stage 3: Final filter
	go finalFilter(outCh)

	// When producers finish, close inCh
	go func() {
		wgProducers.Wait()
		close(inCh)
	}()

	// When processors finish, close outCh
	go func() {
		wgProcessors.Wait()
		close(outCh)
	}()

	// Wait for all to complete
	time.Sleep(10 * time.Second)
}

func producer(name string, start, end, step int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i += step {
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		out <- i
	}
}

func processor(name string, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		out <- val * val
	}
}

func finalFilter(in <-chan int) {
	var lastPrinted *int
	for val := range in {
		if lastPrinted == nil || val > *lastPrinted {
			fmt.Println(val)
			lastPrinted = &val
		}
	}
}
