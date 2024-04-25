package main

import (
	"fmt"
	"log"
)

// the number of attendees we need to serve lunch to
const consumerCount = 300

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string, input []chan string, done chan<- struct{}) {
	for _, course := range input {
		fmt.Printf("%s eats %s\n", name, <-course)
	}
	done <- struct{}{}
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string, output chan<- string, done <-chan struct{}) {
	for {
		select {
		case output <- course:
		case <-done:
			return
		}
	}
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)
	// doneEating channel
	doneEating := make(chan struct{})
	//doneServing channel
	doneServing := make(chan struct{})
	// lista de canais de courses
	var courses []chan string

	// encerrar serverLauch
	defer close(doneServing)

	// loop foodCourses
	for _, course := range foodCourses {
		// criar canal
		ch := make(chan string)
		// colocar canal criado na lista
		courses = append(courses, ch)
		// chamar serveLunch em background
		go serveLunch(course, ch, doneServing)
	}

	// loop consumerCount colocando os Attendees na fila
	for i := 0; i < consumerCount; i++ {
		// variavel name = "Attendee 1"
		name := fmt.Sprintf("Attendee %d", i)
		// chamar takeLunch em background
		go takeLunch(name, courses, doneEating)
	}
	// loop consumerCount aguardando o print na tela
	for i := 0; i < consumerCount; i++ {
		<-doneEating
	}
}
