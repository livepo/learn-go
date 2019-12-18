package main

import "fmt"


type Stack struct {
	list [][]int
}


func NewStack() *Stack {
	list := make([][]int, 0)
	return &Stack{list}
}


func (s *Stack) Push(value []int) {
	s.list = append(s.list, value)
}


func (s *Stack) Top() []int {
	lenl := s.Len()
	if lenl == 0 {
		return nil
	} else {
		return s.list[len(s.list)-1]
	}
}


func (s *Stack) Pop() []int {
	lenl := s.Len()
	if lenl == 0 { return nil }
	value := s.list[lenl-1]
	s.list = s.list[:lenl-1]
	return value
}


func (s *Stack) Len() int {
	return len(s.list)
}


func dailyTemperatures(T []int) []int {
	miniStack := NewStack()
	lent := len(T)
	ret := make([]int, lent)
	for i := 0; i < lent; i ++ { ret[i] = 0 }

	for idx, value := range T {
		// 计算小于今天温度的天数
		for miniStack.Len() > 0 && miniStack.Top()[0] < value {
			v := miniStack.Pop()
			ret[v[1]] = idx - v[1]
		}
		// 将今天温度入最小栈
		miniStack.Push([]int{value, idx})
	}
	for miniStack.Len() > 0 {
		v := miniStack.Pop()
		ret[v[1]] = 0
	}
	return ret
}


func main() {
	// arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	arr := []int{1}
	next_high := dailyTemperatures(arr)
	fmt.Println(next_high)
}

