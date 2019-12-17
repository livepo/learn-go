package main

import "fmt"


func test_2d_array() [][]int {
	arr := make([][]int, 0)
	d1 := []int{1, 2, 3}
	d2 := []int{2, 3, 4}
	d3 := []int{3, 4, 5}
	arr = append(arr, d1)
	arr = append(arr, d2)
	arr = append(arr, d3)
	fmt.Println("arr", arr)
	return arr
}


func main() {
	arr := test_2d_array()
	fmt.Println(arr)
}

