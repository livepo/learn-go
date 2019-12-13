package main


import (
	"fmt"
)


func sum(array [3]int) int {
	s := 0
	for _, v := range array {
		s += v
	}
	return s
}


func main() {
	arr := [3]int{1, 2, 3}
	c := sum(arr)
	fmt.Println(c)
}
