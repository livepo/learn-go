package main


import (
	"fmt"
	"math"
)


type Point struct {
	x, y float64
}


func (p Point) distance(p2 Point) float64 {
	return math.Hypot(p.x-p2.x, p.y-p2.y)
}



func main() {
	p1, p2 := Point{1, 2}, Point{3, 4}
	fmt.Println(p1.distance(p2))
}
