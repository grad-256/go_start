package main

import "fmt"

// func reciever(c chan int) {
// 	for {
// 		i := <-c
// 		fmt.Println(i)
// 	}
// }

// func reciever(name string, ch chan int) {
// 	for {
// 		i, ok := <-ch
// 		if !ok {
// 			break
// 		}
// 		fmt.Println(name, i)
// 	}
// 	fmt.Println(name + "END")
// }

//channel

// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
// 宣言　操作
func main() {
	// 双方方向
	//var ch1 chan int

	// 受信専用
	//var ch2 <-chan int

	// 送信専用
	//var ch3 chan<- int

	// 読み込み書き込み可能
	// ch1 = make(chan int)

	// ch2 := make(chan int)

	// fmt.Println(cap(ch1))
	// fmt.Println(cap(ch2))

	// バッファサイズを超えるとdeadlockになる
	// 取り出しをするとならない
	// ch3 := make(chan int, 5)
	// fmt.Println(cap(ch3))

	// ch3 <- 1

	// fmt.Println(len(ch3))

	// ch3 <- 2
	// ch3 <- 3

	// 最後から順に取り出される
	// キューの性質 先入れ先出し
	// fmt.Println("len", len(ch3))

	// i := <-ch3
	// fmt.Println(i)
	// fmt.Println("len", len(ch3))
	// i2 := <-ch3
	// fmt.Println(i2)
	// fmt.Println("len", len(ch3))

	// fmt.Println(<-ch3)
	// fmt.Println("len", len(ch3))

	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// fmt.Println(<-ch1)
	// go reciever(ch1)
	// go reciever(ch2)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	ch2 <- i
	// 	time.Sleep(50 * time.Millisecond)
	// 	i++
	// }

	// ch1 := make(chan int, 2)

	/*
		ch1 <- 1
		close(ch1)

		// fmt.Println(<-ch1)

		// okのところはチャネルのバッファ内が空且つ開閉状態を表す
		i, ok := <-ch1
		fmt.Println(i, ok)

		i2, ok := <-ch1
		fmt.Println(i2, ok)
	*/
	// go reciever("1.goroutin", ch1)
	// go reciever("2.goroutin", ch1)
	// go reciever("3.goroutin", ch1)
	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	i++
	// }
	// close(ch1)
	// time.Sleep(3 * time.Second)

	// ch1 := make(chan int, 3)
	// ch1 <- 1
	// ch1 <- 2
	// ch1 <- 3
	// close(ch1)
	// for i := range ch1 {
	// 	fmt.Println(i)
	// }

	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch2 <- "A"
	// ch1 <- 1
	// ch2 <- "B"
	// ch1 <- 2
	/*
		v1 := <-ch1
		v2 := <-ch2
		fmt.Println(v1)
		fmt.Println(v2)
	*/

	// チャネルに対する処理である必要がある select
	// deadlockに対する対策で使う
	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("recieved", i3)
		default:
			if n > 100 {
				break L
			}

		}
		// if n > 100 {
		// 	break
		// }
	}
}
