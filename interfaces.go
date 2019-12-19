package main

import "fmt"


type People interface {
	GetName() string
	SetName(name string)
}


type Student struct {
	Name string
	Age int
}


func (s *Student) GetName() string {
	return s.Name
}


func (s *Student) SetName(name string) {
	s.Name = name
}


func main() {
	s := &Student{Name: "nobody", Age: 18}

	var p People
	p = s
	fmt.Println("p", p)
	// fmt.Println(p.GetName())
	s.SetName("Wang")
	fmt.Println("s", s)
	// p = s
	fmt.Println("p", p)
}
