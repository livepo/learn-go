package main


import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	defer close(ch)

	go func() {
		for i := 0; i < 10; i ++ {
			ch <- i
		}
	}()
	for val := range ch {
		fmt.Println(val)
	}


	time.Sleep(time.Second)
}
