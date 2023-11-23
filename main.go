package main

import "fmt"

// const secret = "abc"

type Os int

// constで連番を作ることができる
const (
	Mac     Os = iota + 1 //1
	Windows               //2
	Linux                 //3
)

// var (
// 	i int
// 	s string
// 	b bool
// )

// import (
// 	"fmt"
// 	"go-basics/calculator"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	godotenv.Load()
// 	fmt.Println(os.Getenv("GO_ENV"))
// 	fmt.Println(calculator.Offset)
// 	fmt.Println(calculator.Sum(1, 2))
// 	fmt.Println(calculator.Multiply(1, 2))
// }

func main() {
	//var i int
	//var i int = 2
	// var i = 2
	i := 3
	ui := uint16(2)
	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	f := 1.23456
	s := "buenas tardes"
	b := true
	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("f: %[1]v %[1]T\n", s)
	fmt.Printf("f: %[1]v %[1]T\n", b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println(z)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)

	i = 2
	fmt.Printf("i: %v\n", i)
	i += 1
	fmt.Printf("i: %v\n", i)
	i *= 2
	fmt.Printf("i: %v\n", i)
}
