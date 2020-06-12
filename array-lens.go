package main


import "fmt"


func foo(arr []int) int {
	fmt.Println("foo", len(arr))
	return len(arr)
}



func main() {
	arr := []int{1,2,3,4}
	b := arr[1:]
	fmt.Println("arr", arr)
	fmt.Println("b", b)
	fmt.Println("lenarr, lenb", len(arr), len(b))
	foo(arr[2:])


}
