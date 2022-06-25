package main

import (
	"bufio"
	"net"
	"strings"
)

type ClickerClient struct {
	socket net.Conn
}

func (c ClickerClient) connect(addr string) {
	socket, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	} else {
		c.socket = socket
	}
}

func (c ClickerClient) close() {
	c.socket.Close()
}

func (c ClickerClient) query(q string) string {
	if !strings.HasSuffix(q, ";") {
		q += ";"
	}

	c.socket.Write([]byte(q))
	buff, err := bufio.NewReader(c.socket).ReadString('\n')
	if err != nil {
		panic(err)
	}

	return string(buff)
}
