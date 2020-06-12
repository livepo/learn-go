package main

import (
	"fmt"
	"sort"
)


type IntSlice []int


func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(a, b int) { s[a], s[b] = s[b], s[a] }
func (s IntSlice) Less(a, b int) bool { return s[a] < s[b] }


// [-4, -1, -1, 0, 1, 2]

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	sort.Sort(IntSlice(nums)) // 这里的IntSlice转换不是很清楚

	lenn := len(nums)
	for i := 0; i < lenn-2; i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, lenn-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ret = append(ret, IntSlice{nums[i], nums[j], nums[k]})
				j ++
				for j < k && nums[j] == nums[j-1] {
					j ++
				}
				k --
				for j < k && nums[k] == nums[k+1] {
					k --
				}
			} else if  sum > 0 {
				k --
			} else {
				j ++
			}
		}
	}
	return ret
}


func main() {
	ret := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(ret)
}
