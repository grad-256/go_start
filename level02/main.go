package main

import "fmt"

//変数

/*
明示的な定義なら可能
*/
// i5 := 500
var i5 = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	/*
		明示的な定義
		var 変数名 型 = 値
	*/
	var i int = 100
	fmt.Println(i)

	var s string = "Hellow Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)

	fmt.Println(i2, s2)

	var (
		i3 int
		s3 string
	)
	/*
		0になる から文字が入る
		変数だけの場合は初期値が入る
	*/
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	/*
		値の更新
	*/
	i = 150
	fmt.Println(i)

	/*
		暗黙的な定義
		型指定なしが可能
		関数内のみ可能
		変数名 := 値
	*/
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 := 500
	// fmt.Println(i4)

	/*
		型推論
	*/
	// i4 = "Hellow"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()
}
