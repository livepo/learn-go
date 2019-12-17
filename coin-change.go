package main

import "fmt"



func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}


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



func coinChange(coins []int, amount int) int {
	bubbleSort(coins)
	DP := make([]int, amount+1)
	for idx, _ := range DP {
		DP[idx] = 100
	}
	DP[0] = 0

	for i := 0; i <= amount; i ++ {
		for _, c := range coins {
			if i + c > amount {
				break
			}
			DP[i+c] = Min(DP[i+c], DP[i] + 1)
		}
	}
	if DP[amount] == 100 {
		return -1
	} else {
		return DP[amount]
	}
}


func main() {
	// coins := []int{1, 2, 5}
	coins := []int{474, 83, 404, 3}
	ret := coinChange(coins, 264)
	fmt.Println(ret)
}
