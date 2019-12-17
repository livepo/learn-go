package main


import "fmt"


func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}


func Min(a, b int) int {
	return a + b - Max(a, b)
}


func maxArea(height []int) int {
	lenh := len(height)
	if lenh <= 1 { return 0 }

	l, r := 0, lenh-1
	max := 0
	for l < r {
		area := (r - l) * Min(height[l], height[r])
		max = Max(max, area)
		if height[l] <= height[r] {
			l ++
		} else {
			r --
		}
	}
	return max
}


func main() {
	arr := []int{1,8,6,2,5,4,8,3,7}
	ret := maxArea(arr)
	fmt.Println(ret)
}
