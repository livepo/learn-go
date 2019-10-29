package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "buffered"

	go func(s string) {
		fmt.Println(s)
	}(<-messages)
}
