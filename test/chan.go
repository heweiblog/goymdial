package main

import (
	"fmt"
)

func f1(ch chan bool) {
	fmt.Println("I'm f1 wait a chan data and over myself")
	<-ch
}

func f2(ch chan int, cc chan bool) {
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("recv over")
	<-cc
}

func main() {
	c := make(chan bool)
	go f1(c)
	c <- true
	fmt.Println("send a chan data to f1")
	close(c)
	fmt.Println("f1 over")

	c1 := make(chan int)
	c2 := make(chan bool)
	go f2(c1, c2)
	for i := 1; i < 4; i++ {
		c1 <- i
	}
	close(c1)

	c2 <- true
	close(c2)
	fmt.Println("f2 over")
}
