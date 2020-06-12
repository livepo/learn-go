package main

import "container/list"
import "fmt"


func main() {
	l := list.New()
	for i := 0; i < 10; i ++ {
		l.PushBack(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
