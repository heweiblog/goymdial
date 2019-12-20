package main

import (
	"fmt"
	"sync"
	"time"
)

func say(c chan bool) {
	fmt.Println("say hello")
	fmt.Println("say hello over")
	time.Sleep(time.Second)
	c <- true
}

func say1(w *sync.WaitGroup) {
	fmt.Println("say 1")
	fmt.Println("say 1 over")
	w.Done()
}

func numbers(w *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("say numbers over")
	w.Done()
}

func alphabets(w *sync.WaitGroup) {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println("say char over")
	w.Done()
}

func cal2(num int, c chan int) {
	sum := 0
	for num != 0 {
		n := num % 10
		sum += n * n
		num = num / 10
	}
	c <- sum
}

func cal3(num int, c chan int) {
	sum := 0
	for num != 0 {
		n := num % 10
		sum += n * n * n
		num = num / 10
	}
	time.Sleep(time.Second)
	c <- sum
}

func main() {
	w := new(sync.WaitGroup)
	go numbers(w)
	go alphabets(w)
	go say1(w)
	w.Add(3)
	w.Wait()

	c := make(chan bool)
	go say(c)
	<-c

	//num := 123
	num := 222
	c1 := make(chan int)
	c2 := make(chan int)
	go cal2(num, c1)
	go cal3(num, c2)
	sum := <-c1 + <-c2
	fmt.Println(num, sum)

	fmt.Println("main terminated")
}
