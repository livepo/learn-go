package main

import "fmt"
import "sort"


type IntSlice []int


func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }



func main() {
	a := []int{4,3,2,1,5,9,7,8,6}
	sort.Sort(IntSlice(a))
	fmt.Println("after sorted:", a)
}
