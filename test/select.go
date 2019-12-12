package main

import (
	"fmt"
	"time"
)

func service1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from service1"
}
func service2(ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- "from service2"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go service1(ch1)
	go service2(ch2)
	ok := false
	for {
		select { // 会发送阻塞
		case s1 := <-ch1:
			fmt.Println(s1)
		case s2 := <-ch2:
			fmt.Println(s2)
		case <-time.After(2 * time.Second): // 等待 2s
			fmt.Println("no case ok")
			ok = true
			break
		}
		if ok {
			break
		}
	}
}
