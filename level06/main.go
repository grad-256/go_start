package main

import "fmt"

// 関数

// func Plus(x, y int) int {
// 	return x + y
// }

// func Div(x, y int) (int, int) {
// 	q := x / y
// 	r := x % y
// 	return q, r
// }

// // 返り値の変数を指定
// func Double(price int) (result int) {
// 	result = price * 2
// 	return result
// }

// func Noreturn() {
// 	fmt.Println("No Return")
// 	return
// }

// func ReturnFunc() func() {
// 	return func() {
// 		fmt.Println("Im a function")
// 	}
// }

// func CallFunction(f func()) {
// 	f()
// }

// func Later() func(string) string {
// 	var store string
// 	return func(next string) string {
// 		s := store
// 		store = next
// 		return s

// 	}
// }

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// i := Plus(1, 2)

	// fmt.Println(i)

	// i2, i3 := Div(9, 3)
	// fmt.Println(i2, i3)

	// i4 := Double(1000)
	// fmt.Println(i4)

	// Noreturn()

	// 無名関数
	// f := func(x, y int) int {
	// 	return x + y
	// }
	// i := f(1, 2)
	// fmt.Println(i)

	// i2 := func(x, y int) int {
	// 	return x + y
	// }(1, 2)
	// fmt.Println(i2)

	// 関数を返す関数
	// f := ReturnFunc()
	// f()

	//関数を引数に取る関数
	// CallFunction(func() { fmt.Println("Im a function") })

	// クロージャ
	// f := Later()
	// fmt.Println(f("Hello"))
	// fmt.Println(f("my"))
	// fmt.Println(f("name"))
	// fmt.Println(f("is"))
	// fmt.Println(f("Golang"))

	// ジェネレーター ルールに従って連続した値を返し続ける
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}
