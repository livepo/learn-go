package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"./network"
)

var (
	localServerAddr = "127.0.0.1:8000"
	remoteIP = "livepo.info"
	remoteControlAddr = remoteIP + ":8009"
	remoteServerAddr = remoteIP + ":8008"
)

func main() {
	tcpConn, err := network.CreateTCPConn(remoteControlAddr)
	if err != nil {
		log.Println("[连接失败]" + remoteControlAddr + err.Error())
		return
	}
	log.Println("[已连接]" + remoteControlAddr)

	reader := bufio.NewReader(tcpConn)

	for {
		s, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}

		if s == network.NewConnection + "\n" {
			go connectLocalAndRemote()
		}
	}
	log.Println("[已断开]" + remoteControlAddr)
}


func connectLocalAndRemote() {
	local := connectLocal()
	remote := connectRemote()

	if local != nil && remote != nil {
		network.Join2Conn(local, remote)
	} else {
		if local != nil {
			_ = local.Close()
		}
		if remote != nil {
			_ = remote.Close()
		}
	}
}


func connectLocal() *net.TCPConn {
	conn, err := network.CreateTCPConn(localServerAddr)
	if err != nil {
		log.Println("[连接本地服务失败]" + err.Error())
	}
	return conn
}


func connectRemote() *net.TCPConn {
	conn, err := network.CreateTCPConn(remoteServerAddr)
	if err != nil {
		log.Println("[连接远端服务失败]" + err.Error())
	}
	return conn
}
