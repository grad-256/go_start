package main

import (
	"fmt"
	"sqlite/alib"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(2))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
