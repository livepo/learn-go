package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
        out := x
	const lim float64 = 0.00001
	for ;; {
                if out = z - (z*z - x)/(2*z); (out-z <= lim) && (out-z >= -1*lim) {
			return out
		} else {
			z = out
		}

	}
}

func main() {
	fmt.Println(Sqrt(2))
}
