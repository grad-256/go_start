package main

import "fmt"

// interface
// システムに柔軟性を持たせられる
// 型が違った場合にまとめられない時もまとめられる

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v , Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (p *Car) ToString() string {
	return fmt.Sprintf("Number=%v , Model=%v", p.Number, p.Model)
}

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	p := Point{100, "ABC"}
	fmt.Println(p)
	// ポインタを通してアクセスする
	pp := &Point{100, "ABC"}
	fmt.Println(pp)
}
