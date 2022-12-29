package main

import "fmt"

//定数

const Pi = 3.14

const (
	URL     = "http://***.co.jp"
	SITNAME = "***"
)

// iotaは連番を返す
const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)
	fmt.Println(URL, SITNAME)

	fmt.Println(c0, c1, c2)
}
