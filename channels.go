package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	go func() {
		messages <- "ping1"
		messages <- "ping2"
		messages <- "ping3"
	}()

	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}
