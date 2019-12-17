package main

import "fmt"


func rotate(matrix [][]int)  {
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 { return }

	for r := 1; r < row; r ++ {
		for c := 0; c < r; c ++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	for r := 0; r < row; r ++ {
		for c := 0; c < col/2; c++ {
			matrix[r][c], matrix[r][col-c-1] = matrix[r][col-c-1], matrix[r][c]
		}
	}
}


func main() {
	m := make([][]int, 0)
	m = append(m, []int{1,2,3})
	m = append(m, []int{4,5,6})
	m = append(m, []int{7,8,9})
	rotate(m)
	fmt.Println(m)
}
