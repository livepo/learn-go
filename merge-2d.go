package main


import (
	"fmt"
	"sort"
)


type Slice [][]int

func (s Slice) Swap(idxa, idxb int) {
    s[idxa], s[idxb] = s[idxb], s[idxa]
}

func (s Slice) Len() int {
    return len(s)
}

func (s Slice) Less(idxa, idxb int) bool {
    if s[idxa][0] < s[idxb][0] { return true }
    if s[idxa][0] == s[idxb][0] && s[idxa][1] <= s[idxb][1] { return true }
    return false
}

func merge(intervals [][]int) [][]int {
    len_intervals := len(intervals)

    if len_intervals == 1 {
        return intervals
    } else {
        sort.Sort(Slice(intervals))
        return reduce_merge(intervals, 0)
    }
}

func reduce_merge(intervals [][]int, pos int) [][]int {
	fmt.Println("pos", pos)
    len_intervals := len(intervals)

    if pos >= len_intervals-1 { return intervals }
    first := intervals[pos]
    second := intervals[pos+1]
    if first[1] >= second[0] {
        merged := []int{first[0], Max(first[1], second[1])}
        concat := [][]int{}
        if pos-1>= 0 {
            concat = append(concat, intervals[:pos]...)
        }
        concat = append(concat, merged)
        if pos+2 <= len_intervals {
            concat = append(concat, intervals[pos+2:]...)
        }
        return reduce_merge(concat, pos)
    } else {
        return reduce_merge(intervals, pos+1)
    }

}

func Max(a, b int) int {
    if a <= b {
        return b
    } else {
        return a
    }
}


func main() {
	intervals := [][]int{
		[]int{1, 3},
		[]int{2, 2},
		[]int{2, 3},
		[]int{4, 6},
		[]int{5, 7},
	}
	result := merge(intervals)
	fmt.Println(result)
}
