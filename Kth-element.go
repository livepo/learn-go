package main

import "fmt"


type Heap struct {
	Array []int
	K int
}


func NewHeap(K int) Heap {
	arr := make([]int, 0)
	return Heap{Array: arr, K: K}
}


func (h *Heap) Insert(value int) {
	fmt.Println("befor", h.Array)
	lena := len(h.Array)
	if lena < h.K {
		h.Array = append([]int{value}, h.Array...)
		h.Heapify(0)
	} else {
		if value > h.Top() {
			h.Array[0] = value
			h.Heapify(0)
		}
	}
	fmt.Println("after", h.Array)
}


func (h *Heap) Top() int {
	return h.Array[0]
}


func (h *Heap) Heapify(pos int) {
	lenh := len(h.Array)
	last_node := (lenh-1) / 2
	if pos > last_node { return }
	left, right := 2*pos+1, 2*pos+2
	max := pos
	if left < lenh && h.Array[pos] > h.Array[left] { max = left }
	if right < lenh && h.Array[max] > h.Array[right] { max = right }  // 注意这里是max跟right比较,关键的腾挪操作!
	if max != pos {
		h.Array[max], h.Array[pos] = h.Array[pos], h.Array[max]
		h.Heapify(max)
	}
}


func findKthLargest(nums []int, k int) int {
	heap := NewHeap(k)
	for _, n := range nums {
		heap.Insert(n)
	}
	return heap.Top()
}


func main() {
	arr := []int{3,2,3,1,2,4,5,5,6}
	ret := findKthLargest(arr, 4)
	fmt.Println(ret)
}
