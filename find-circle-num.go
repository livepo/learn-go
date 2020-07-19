package main

import "fmt"


func findCircleNum(M [][]int) int {
	row := len(M)
	col := len(M[0])
	if row == 0 || col == 0 { return 0 }

	count := 0
	for i := 0; i < row; i ++ {
		for j := 0; j < col; j ++ {
			if M[i][j] == 1 {
				count += 1
				DFS(M, i, j)
			}
		}
	}

	for i := 0; i < row; i ++ {
		for j := 0; j < col; j ++ {
			if M[i][j] == 2 {
				count += 1
			}
		}
	}

	return count
}


func DFS(M [][]int, i int, j int) {
	row, col := len(M), len(M[0])
	if i < 0 || i >= row || j < 0 || j >= col || M[i][j] == 0 { return }
	M[i][j] = 0
	DFS(M, i+1, j)
	DFS(M, i-1, j)
	DFS(M, i, j+1)
	DFS(M, i, j-1)
}


func main() {
	M := [][]int{
		[]int{1,0,0,1},
		[]int{0,1,1,0},
		[]int{0,1,1,1},
		[]int{1,0,1,1},
	}
	fmt.Println(findCircleNum(M))
}
