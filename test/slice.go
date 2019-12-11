package main

import (
	"fmt"
)

func changeSlice1(s []int) {
	s = append(s, 1, 2, 3)
}

func changeSlice2(s *[]int) {
	*s = append(*s, 1, 2, 3)
}

func changeSlice3(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}
}

func changeMap(s map[string]int) {
	s["1"] = 1
	s["2"] = 2
}

func main() {
	s := make([]int, 0)
	changeSlice1(s)
	fmt.Println(s)
	changeSlice2(&s)
	fmt.Println(s)
	changeSlice3(s)
	fmt.Println(s)

	m := make(map[string]int)

	changeMap(m)
	fmt.Println(m)

	if m["3"] == 0 {
		fmt.Println("m[3] =", m["3"])
	}
	if a, ok := m["3"]; ok {
		fmt.Println("m[3] =", a)
	} else {
		fmt.Println("m[3] not exist")
	}
}
