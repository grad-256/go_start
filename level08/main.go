package main

import "fmt"

// func Sum(s ...int) int {
// 	n := 0
// 	for _, v := range s {
// 		n += v
// 	}
// 	return n
// }

func main() {
	// var sl []int
	// fmt.Println(sl)

	// var sl2 []int = []int{100, 200}
	// fmt.Println(sl2)

	// sl3 := make([]int, 5)
	// fmt.Println(sl3)
	// // fmt.Printf("%T\n", sl3)

	// sl2[0] = 1000
	// fmt.Println(sl2)

	// sl5 := []int{1, 2, 3, 4, 5}
	// fmt.Println(sl5)

	// fmt.Println(sl5[0])

	// fmt.Println(sl5[2:4])

	// fmt.Println(sl5[:2])

	// fmt.Println(sl5[2:])

	// fmt.Println(sl5[:])

	// fmt.Println(sl5[len(sl5)-1])

	// fmt.Println(sl5[1 : len(sl5)-1])

	// スライス append make ken cap copy
	/*
		cap
		パフォーマンス面で影響が出る
		キャパシティを気にする開発では必要
	*/
	// sl := []int{100, 200}
	// fmt.Println(sl)

	// sl = append(sl, 300)
	// fmt.Println(sl)

	// sl = append(sl, 400, 500, 600)
	// fmt.Println(sl)

	// sl2 := make([]int, 5)
	// fmt.Println(sl2)

	// fmt.Println(len(sl2))
	// fmt.Println(cap(sl2))

	// fmt.Println("len & cap")
	// sl3 := make([]int, 5, 10)
	// fmt.Println(sl3)
	// fmt.Println(len(sl3))
	// fmt.Println(cap(sl3))
	// fmt.Printf("len=%d cap=%d value=%v\n", len(sl3), cap(sl3), sl3)

	// sl := []int{100, 200}
	// sl2 := sl
	// sl2[0] = 1000

	// fmt.Println(sl)

	// var i int = 10
	// i2 := i
	// i2 = 100
	// fmt.Println(i, i2)

	// sl := []int{1, 2, 3, 4, 5}
	// sl2 := make([]int, 5, 10)
	// fmt.Println(sl2)

	// n := copy(sl2, sl)

	// fmt.Println(n, sl2)

	// fmt.Println(Sum(1, 2, 3))
	// fmt.Println(Sum())

	// sl := []int{1, 2, 3}
	// fmt.Println(Sum(sl...))

	var m = map[string]int{"A": 100, "B": 200}

	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{1: "A", 2: "B"}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3]) // ないので初期値が入る

	s, ok := m4[5]
	if !ok {
		fmt.Println("error")
	} else {
		fmt.Println(s)
	}

	fmt.Println(s, ok)

	m4[3] = "CHINE"
	fmt.Println(m4)
	delete(m4, 3)
	fmt.Println(m4)

	m5 := map[int][]string{
		1: {"A"},
		2: {"B", "C"},
		3: {"D", "E"},
	}
	fmt.Println(m5)
}
