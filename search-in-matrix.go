package main


import "fmt"


func findNumberIn2DArray(matrix [][]int, target int) bool {
    row := len(matrix)
    if row == 0 { return false }

    col := len(matrix[0])
    if col == 0 { return false }

    fmt.Println("row, col", row, col)
    i, j := 0, row-1
    for i < j {
        mid := (i+j)/2
        if matrix[mid][col-1] == target {
            return true
        } else if matrix[mid][col-1] > target {
            j = mid-1
        } else {
            i = mid+1
        }
    }
    fmt.Println("i,j", i, j)

    k := 0
    j = col-1
    for k < j {
        mid := (k+j)/2
        if matrix[i][mid] == target {
            return true
        } else if matrix[i][mid] > target {
            j = mid-1
        } else {
            k = mid+1
        }
    }
    fmt.Println("k,j", k, j)
    return false
}


func main() {
	matrix := make([][]int, 0)
	matrix = append(matrix, []int{1,4,7,11,25})
	matrix = append(matrix, []int{2,5,8,12,19})
	matrix = append(matrix, []int{3,6,9,16,22})
	matrix = append(matrix, []int{10,13,14,17,24})
	matrix = append(matrix, []int{18,21,23,26,30})
	fmt.Println(findNumberIn2DArray(matrix, 5))
}
