package main

import (
	"fmt"
	"strconv"
)

// ジェネリクス
// 中身は一緒でも方が違う場合でも
// コードの再利用できるようにする
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type Stringer interface {
	String() string
}

func f[T Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

type Vector[T any] []T

type IntVector = Vector[int]

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

type Set[T comparable] map[T]struct {
}

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, v := range xs {
		s.Add(v)
	}

	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func main() {
	// PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// PrintSlice[string]([]string{"a", "b", "c"})

	// fmt.Println(f([]MyInt{1, 2, 4, 3}))

	// fmt.Println(Max(1, 2))
	// fmt.Println(Max(1.1, 2.5))

	// var x, y MyInt = 1, 2
	// fmt.Println(Max[MyInt](x, y))

	// var v Vector[int] = Vector[int]{10, 20, 30}
	// fmt.Println(v)

	// var v2 Vector[float64] = Vector[float64]{1.3, 3.4}
	// fmt.Println(v2)

	// v3 := IntVector{1, 2, 3}
	// fmt.Println(v3)

	// var t T[int, []*int, *int]
	// fmt.Println("A: %T, B: %T, C: %T\n", t.a, t.b, t.c)

	s := NewSet(1, 2, 3)
	fmt.Println(s)

	s.Remove(1)
	fmt.Println(s)
}
