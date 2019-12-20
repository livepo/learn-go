package main

import "fmt"


func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	locate_r, locate_c := 0, 0

	for l, r := 0, row-1; l < r; {
		mid := (l + r) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			r = mid
		} else {
			l = mid
		}




}


func main() {
	arr := make([][]int, 0)
	arr = append(arr, []int{1,   4,  7, 11, 15})
	arr = append(arr, []int{2,   5,  8, 12, 19})
	arr = append(arr, []int{3,   6,  9, 16, 22})
	arr = append(arr, []int{10, 13, 14, 17, 24})
	arr = append(arr, []int{18, 21, 23, 26, 30})
	ret = searchMatrix(arr, 5)
	fmt.Println(ret)
}
