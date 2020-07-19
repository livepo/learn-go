package main

import "fmt"


func findCircleNum(M [][]int) int {
	row := len(M)
	if row == 0 { return 0 }
	col := len(M[0])
	if col == 0 { return 0 }

	Q := make([]int, len(M))
	for i := 0; i < row; i ++ {
		Q[i] = i
	}

	for i := 0; i < len(M); i ++ {
		for j := 0; j < len(M[0]); j ++ {
			if i == j { continue }
			if M[i][j] == 1 {
				if i > j {
					Q[i] = findRoot(j)
				} else {
					Q[j] = findRoot(i)
				}
			}
		}
	}

	countMap := make(map[int]struct{}, 0)
	for i := 0; i < row; i ++ {
		x := i
		for Q[x] != x {
			x = Q[x]
		}
		countMap[x] = struct{}{}
	}
	return len(countMap)
}


func findRoot(Q []int, i int) int {
	if Q[i] == i {
		return i
	} else {
		return findRoot(Q, Q[i])
	}
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
