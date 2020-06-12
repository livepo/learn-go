package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("conn server failed, err:%v\n", err)
		return
	}

	input := bufio.NewReader(os.Stdin)

	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)

		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err = conn.Write([]byte(s))

		if err != nil {
			fmt.Println("send failed, err: %v\n", err)
			return
		}

		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
			return
		}
		fmt.Printf("收到服务端回复：%v\n", string(buf[:n]))
	}
}
