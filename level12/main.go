package main

import (
	"fmt"
	"os"
)

type Dog struct{}

func (d *Dog) Bark() string {
	return "Bow"
}

type BullDog struct{ Dog }
type ShiobaInu struct{ Dog }

func (d *ShiobaInu) Bark() string { return "ワン" }
func DogVoice(d *Dog) string {
	return d.Bark()
}

type Warning interface {
	Show(message string)
}

type ConsoleWarning struct{}

func (c ConsoleWarning) Show(message string) {
	fmt.Fprintf(os.Stderr, "[%s]: %s\n", os.Args[0], message)
}

type DeskTopWarning struct{}

func main() {
	var warn Warning = &ConsoleWarning{}

	// warn =

	warn.Show("Hellow World to console")

	// bd := &BullDog{}
	// fmt.Println(bd.Bark())
	// si := &ShiobaInu{}
	// fmt.Println(si.Bark())

	// var bg *string

	// g := "test"

	// bg = &g

	// fmt.Println(*bg)

}
