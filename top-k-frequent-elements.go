package main

import "fmt"
import "sort"


// 返回频次最高的k个元素
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	ret := make([]int, 0)
	f := make([]int, 0)
	for _, v := range nums {
		m[v] ++
	}
	for _, v := range m {
		f = append(f, v)
	}
	sort.Ints(f)
	pivot := f[len(f)-k]
	for k, v := range m {
		if v >= pivot {
			ret = append(ret, k)
		}
	}


	return ret
}


func main() {
	nums := []int{1,2}
	ret := topKFrequent(nums, 2)
	fmt.Println(ret)
}
