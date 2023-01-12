package main

import "fmt"

// type Dog struct{}

// func (d *Dog) Bark() string {
// 	return "Bow"
// }

// type BullDog struct{ Dog }
// type ShiobaInu struct{ Dog }

// func (d *ShiobaInu) Bark() string { return "ワン" }
// func DogVoice(d *Dog) string {
// 	return d.Bark()
// }

// type Warning interface {
// 	Show(message string)
// }

// type ConsoleWarning struct{}

// func (c ConsoleWarning) Show(message string) {
// 	fmt.Fprintf(os.Stderr, "[%s]: %s\n", os.Args[0], message)
// }

// type DeskTopWarning struct{}

func makeCh() chan int {
	return make(chan int)
}

// <-chan int 受信用　チャネル
func recvCh(recv <-chan int) int {
	return <-recv
}

func main() {
	// var warn Warning = &ConsoleWarning{}

	// warn =

	// warn.Show("Hellow World to console")

	// bd := &BullDog{}
	// fmt.Println(bd.Bark())
	// si := &ShiobaInu{}
	// fmt.Println(si.Bark())

	// var bg *string

	// g := "test"

	// bg = &g

	// fmt.Println(*bg)

	// ch := make(chan int)
	// go func() {
	// 	ch <- 100
	// 	// fmt.Println("別のゴールーチン")
	// }()
	// go func() {
	// 	v := <-ch
	// 	fmt.Println(v)
	// 	// fmt.Println("別のゴールーチン")
	// }()
	// time.Sleep(2 * time.Second)
	// fmt.Println("mainゴールーチン")

	// ch1 := make(chan int)
	// var ch2 chan string
	// go func() { ch1 <- 100 }()
	// go func() { ch2 <- "hi" }()

	// select {
	// case v1 := <-ch1:
	// 	fmt.Println(v1)
	// case v2 := <-ch2:
	// 	fmt.Println(v2)
	// }
	ch := makeCh()
	// chan<- int 送信用チャネル
	go func(ch chan<- int) { ch <- 100 }(ch)
	fmt.Println(recvCh(ch))

}
