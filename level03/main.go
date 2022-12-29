package main

import "fmt"

/*
型
整数型
*/
func main() {
	// var i int = 100

	/*
		int環境依存
		マシンによって異なる

		int8, int16, int32, int64
	*/

	// var i2 int64 = 200

	// fmt.Println(i + 50)

	// ミスマッチが起きる
	// fmt.Println(i + i2)

	/*
		型変換
	*/
	// 現在の型を調べる
	// fmt.Printf("%T\n", i)
	// fmt.Printf("%T\n", i2)

	// fmt.Printf("%T\n", int64(i2))

	// fmt.Println(i + int(i2))

	// byteA := []byte{72, 73}
	// fmt.Println(byteA)

	// fmt.Println(string(byteA))

	// c := []byte("HI")
	// fmt.Println(c)

	// fmt.Println(string(c))

	// var arr1 [3]int
	// fmt.Println(arr1)
	// fmt.Printf("%T\n", arr1)

	// var arr2 [3]string = [3]string{"A", "B"}
	// fmt.Println(arr2)

	// arr3 := [3]int{1, 2, 3}
	// fmt.Println(arr3)

	// arr4 := [...]string{"C", "D"}
	// fmt.Println(arr4)
	// fmt.Printf("%T\n", arr4)

	// fmt.Println(arr2[0])
	// fmt.Println(arr2[1])
	// fmt.Println(arr2[2])
	// fmt.Println(arr2[2-1])

	// arr2[2] = "C"
	// fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// fmt.Println(len(arr1))

	// var slices []string = []string{"e", "r", "rr", "kkk"}
	// fmt.Printf("%T\n", slices)
	// fmt.Println(len(slices))
	// fmt.Println(slices)

	/*
		初期値はnil
		あらゆる方の値を入れられる
		interfaceと計算はできない
		汎用的に使うためにあるものなので、演算には使用しない
	*/
	// var x interface{}
	// fmt.Println(x)

	// x = 1
	// fmt.Println(x)
	// x = "Hellow"
	// fmt.Println(x)
	// x = [3]int{1, 2, 3}
	// fmt.Println(x)

	/*
		型変換
	*/
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)

	// var s string = "100"
	// fmt.Printf("%T\n", s)

	// i, _ := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(i)
	// }
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)

	// var i2 int = 200
	// s2 := strconv.Itoa(i2)
	// fmt.Println(s2)
	// fmt.Printf("s2 = %T\n", s2)

	var h string = "Hellow World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
