package main

import (
	"log"
	"runtime"
	"testing"
	"time"
)

func TestServeLunch(t *testing.T) {
	initialGoroutines := runtime.NumGoroutine()
	doneServing := make(chan struct{})
	ch := make(chan string)
	c := "lanche"
	go serveLunch(c, ch, doneServing)
	v := <-ch
	if v != c {
		t.Fatalf("expected: %s, got: %s\n", c, v)
	}
	doneServing <- struct{}{}
	time.Sleep(100 * time.Millisecond)
	finalGoroutines := runtime.NumGoroutine()
	if finalGoroutines != initialGoroutines {
		t.Fatalf("expected: %d, got: %d\n", initialGoroutines, finalGoroutines)
	}
}

func TestTakeLunch(t *testing.T) {
	var courses []chan string
	doneEating := make(chan struct{})
	doneServing := make(chan struct{})

	ch := make(chan string)
	c := "lanche"
	go serveLunch(c, ch, doneServing)
	courses = append(courses, ch)

	go takeLunch("Inês", courses, doneEating)
	go takeLunch("Celso", courses, doneEating)

	<-doneEating
	<-doneEating

	close(doneServing)
}

func TestTwo(t *testing.T) {
	var channels []chan string
	ch1 := make(chan string)
	ch2 := make(chan string)
	channels = append(channels, ch1)
	go serverAll("Nega1", ch1)
	channels = append(channels, ch2)
	go serverAll("Nega2", ch2)
	go readerOne(channels)
	go readerTwo(channels)
	time.Sleep(100 * time.Millisecond)
}
func serverAll(value string, ch chan<- string) {
	for {
		ch <- value
	}
}
func readerOne(channels []chan string) {
	for _, ch := range channels {
		log.Println("readerOne: ", <-ch)
	}
}
func readerTwo(channels []chan string) {
	for _, ch := range channels {
		log.Println("readerTwo: ", <-ch)
	}
}

// func readerTwo(ch chan string) {
// 	log.Println("readerTwo: ", <-ch)
// }
