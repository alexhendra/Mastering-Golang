package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello Alex"
}

func main() {
	ch := make(chan string)
	go greet(ch)
	ch <- "Hello Genk"

	log.Println(<-ch)
	log.Println(<-ch)
}
