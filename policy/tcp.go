package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//conn, err := net.DialTimeout("tcp", "192.168.6.55:22", time.Second)
	conn, err := net.DialTimeout("tcp6", "[::1]:22", time.Second)
	if err == nil {
		fmt.Println(conn, conn.LocalAddr(), conn.RemoteAddr())
	} else {
		fmt.Println(err)
	}
	conn.Close()
}
