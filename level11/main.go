package main

import "fmt"

// ジェネリクス
// 中身は一緒でも方が違う場合でも
// コードの再利用できるようにする
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice[string]([]string{"a", "b", "c"})
}
