package main

import (
	"fmt"
	"container/list"
)


func match(a, b rune) bool {
	switch(b) {
	case '}':
		return a == '{'
	case ')':
		return a == '('
	case ']':
		return a == '['
	}
	return false
}


func isValid(s string) bool {
	stack := list.New()
	for _, c := range s {
		last := stack.Back()
		if last == nil || !match(last.Value.(rune), c) {
			stack.PushBack(c)
		} else {
			stack.Remove(stack.Back())
		}
	}
	if stack.Len() != 0 {
		return false
	} else {
		return true
	}
}


func main() {
	ret := isValid("({)}()")
	fmt.Println(ret)
}
