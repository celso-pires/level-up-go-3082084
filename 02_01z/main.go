package main

import (
	"flag"
	"log"
	"sync"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeatWithWaitGroup concurrently prints out the given message n times
func repeatWithWaitGroup(n int, message string) {
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			log.Printf("[G%d]: %s\n", i, message)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// repeatWithChannel concurrently prints out the given message n times
func repeatWithChannel(n int, message string) {

	ch := make(chan struct{})
	for i := 0; i < n; i++ {
		go func(i int) {
			log.Printf("[G%d]: %s\n", i, message)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < n; i++ {
		<-ch
	}
}

func main() {
	// const chosenMode = "WithWaitGroup"
	const chosenMode = "WithChannel"
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	for _, m := range messages {
		log.Println(m)
		if chosenMode == "WithChannel" {
			repeatWithChannel(int(*factor), m)
		} else {
			repeatWithWaitGroup(int(*factor), m)
		}
	}
}
