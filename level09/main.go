package main

import "fmt"

//ポインタ
// 値型に分類されるデータ構造 基本型 参照型 構造体
// メモリ上のアドレスと型の情報

func Double(i int) {
	i = i * 2
}

// コピーしないで渡している
func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	// fmt.Println(n)

	// fmt.Println(&n)

	Double(n)
	// fmt.Println(n)
	// &n 参照元アドレス
	// *p 参照先変数

	// 参照渡し
	// *がつくとポインタ型
	// &はアドレス演算子
	var p *int = &n
	// fmt.Println(p)
	// *pは実体
	fmt.Println(p)  //アドレスが出る
	fmt.Println(*p) // 実体が出る 100

	// 置き換わる
	// アドレスを共有できる
	// *p = 300
	// fmt.Println(n)
	// n = 200
	// fmt.Println(*p)
	// ポインタ型は型とメモリー上のアドレスを組み合わせたデータ型で値型

	Doublev2(&n)
	fmt.Println(n)
	// fmt.Println(*p)
	Doublev2(p)
	fmt.Println(*p)

	// 参照型のデータ構造の場合は更新される
	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl)
}
