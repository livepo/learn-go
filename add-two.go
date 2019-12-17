package main


import (
	"fmt"
)


//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			for _, ix := range idx {
				if ix != i {
					return []int{i, ix}
				}
			}
		}
	}
	return []int{-1, -1}
}


func main() {
	nums := []int{2, 5, 5, 11}
	c := twoSum(nums, 10)
	fmt.Println(c)
}
