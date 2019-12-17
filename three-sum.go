package main
import "fmt"


func bubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i ++ {
		for j := i; j < length; j ++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return
}


func threeSum(nums []int) [][]int {
	length := len(nums)
	bubbleSort(nums)
	ret := make([][]int, 0)
	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, length-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j ++
				for j < k && nums[j] == nums[j-1] {
					j ++
				}
				k --
				for j < k && nums[k] == nums[k+1] {
					k --
				}
			} else if sum > 0 {
				k --
			} else {
				j ++
			}
		}
	}
	return ret
}


func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	ret := threeSum(arr)
	fmt.Println(ret)
}
