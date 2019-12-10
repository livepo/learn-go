package main
import (
	"fmt"
	"os"
	"strings"
)


func main1() {
	s, sep := "", ""
	for _, arg := range os.Args[2:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}


func main2() {
	fmt.Println(strings.Join(os.Args[2:], " "))
}


func main() {
	if os.Args[1] == "1" {
	        main1()
	} else {
		main2()
	}
}

