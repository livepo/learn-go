package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Timestamp())
	fmt.Println("My favorite number is", rand.Intn(10))
}
