package main


import "fmt"


type People interface {
}


func main() {
	m := make(map[string]interface{})  // make不需要传那个0进去啊
	m["hello"] = 1
	m["world"] = "golang"
	fmt.Println(m)
}
