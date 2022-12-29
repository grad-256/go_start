package main

import "fmt"

// func TestDefer() {
// 	defer fmt.Println("END")
// 	defer fmt.Println("222")
// 	defer fmt.Println("111")
// 	fmt.Println("START")
// 	defer fmt.Println("333")
// }

// func sub() {
// 	for {
// 		fmt.Println("Sub loop")
// 		time.Sleep(100 + time.Millisecond)
// 	}
// }

// mainより先に実行
func init() {
	fmt.Println("init")
}

func main() {
	// if
	// 条件分岐

	// a := 1
	// if a == 2 {
	// 	fmt.Println("two")
	// } else if a == 1 {
	// 	fmt.Println("one")
	// } else {
	// 	fmt.Println("no")
	// }

	// if b := 100; b == 100 {
	// 	fmt.Println("one hundred")
	// }

	// x := 0
	// if x := 2; true {
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// エラーハンドリング
	// var s string = "S"

	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("i = %T\n", i)

	//for
	//繰り返し処理

	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		fmt.Println("3です")
	// 		continue
	// 	}
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1, 2, 3}
	// // for i := 0; i < len(arr); i++ {
	// // 	fmt.Println(arr[i])
	// // }

	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// sl := []string{"Python", "PHP", "GO"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// m := map[string]int{"apple": 100, "banana": 200}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// n := 0
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("どれでもない")
	// }

	// switch {
	// case n == 0 && n < 4:
	// 	fmt.Println(".....")
	// case n < 3 && n > 2:
	// 	fmt.Println("////")
	// }

	// defer
	// 一番下から遅れて実行
	// TestDefer()

	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// file.Write([]byte("Hello\n"))

	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }()

	// panic("rungime error")
	// fmt.Println("START")

	// panic("runtime error")
	// fmt.Println("START")

	// go sub()
	// // go sub()

	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	fmt.Println("main")
}
