package main


import (
	"fmt"
)


func add2(a, b int) int {
	return a + b
}


func main() {
	c := add2(1, 2)
	fmt.Println(c)
}
