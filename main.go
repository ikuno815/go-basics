package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// // const secret = "abc"

// type Os int

// // constで連番を作ることができる
// const (
// 	Mac     Os = iota + 1 //1
// 	Windows               //2
// 	Linux                 //3
// )

// // var (
// // 	i int
// // 	s string
// // 	b bool
// // )

// // import (
// // 	"fmt"
// // 	"go-basics/calculator"
// // 	"os"

// // 	"github.com/joho/godotenv"
// // )

// // func main() {
// // 	godotenv.Load()
// // 	fmt.Println(os.Getenv("GO_ENV"))
// // 	fmt.Println(calculator.Offset)
// // 	fmt.Println(calculator.Sum(1, 2))
// // 	fmt.Println(calculator.Multiply(1, 2))
// // }

// func main() {
// 	//var i int
// 	//var i int = 2
// 	// var i = 2
// 	i := 3
// 	ui := uint16(2)
// 	fmt.Println(i)
// 	fmt.Printf("i: %v %T\n", i, i)
// 	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

// 	f := 1.23456
// 	s := "buenas tardes"
// 	b := true
// 	fmt.Printf("f: %[1]v %[1]T\n", f)
// 	fmt.Printf("f: %[1]v %[1]T\n", s)
// 	fmt.Printf("f: %[1]v %[1]T\n", b)

// 	pi, title := 3.14, "Go"
// 	fmt.Printf("pi: %v title: %v\n", pi, title)

// 	x := 10
// 	y := 1.23
// 	z := float64(x) + y
// 	fmt.Println(z)

// 	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)

// 	i = 2
// 	fmt.Printf("i: %v\n", i)
// 	i += 1
// 	fmt.Printf("i: %v\n", i)
// 	i *= 2
// 	fmt.Printf("i: %v\n", i)

// 	var ui1 uint16
// 	fmt.Printf("memory address of ui1: %p\n", &ui1)
// 	var ui2 uint16
// 	fmt.Printf("memory address of ui1: %p\n", &ui2)

// 	//ポインタ変数を定義する
// 	//ポインタ変数の型は＊で表現する
// 	var p1 *uint16
// 	fmt.Printf("value of p1: %v\n", p1)

// 	//p1のポインタ変数にui1のアドレスを代入する
// 	p1 = &ui1
// 	fmt.Printf("value of p1: %v\n", p1)

// 	//ポインタ変数のポインタ
// 	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
// 	fmt.Printf("memory address of p1: %v\n", &p1)

// 	//ポインタ変数から、そのアドレスが指し示す値を取得することをdereference*という
// 	fmt.Printf("value of ui1(dereference): %v\n", *p1)
// 	//↑ui1の初期値である1がprintされる
// 	//dereferenceの表記を使って、ui1の値を書き換えることもできる
// 	*p1 = 1
// 	fmt.Printf("value of ui1: %v\n", ui1)

// 	//ダブルポインタpp1を定義する
// 	//p1のポインタ変数の先頭アドレスをampersandで取得してpp1に代入している
// 	var pp1 **uint16 = &p1
// 	fmt.Printf("value of pp1: %v\n", pp1)
// 	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))

// 	//1回、2回ereferenceを使って値を取得する
// 	fmt.Printf("value of p1(dereference): %v\n", *pp1)
// 	fmt.Printf("value of ui1(dereference): %v\n", **pp1)

// 	//2回のdereferenceを使って値を書き換えることもできる
// 	**pp1 = 10
// 	fmt.Printf("value of ui1: %v\n", ui1)

// 	//shadowingについて

// 	//okとresultという変数を用意し、それぞれtrueとAという初期値を代入する
// 	//下記を実行すると、B,Aがprintされる。:=はif文の中のスコープでのみ適用されるから。
// 	// ok, result := true, "A"
// 	// fmt.Printf("memory address of result: %p\n", &result) //119行目と122行目のresultは別物
// 	// if ok {
// 	// 	result := "B"
// 	// 	fmt.Printf("memory address of result: %p\n", &result)
// 	// 	println(result)
// 	// } else {
// 	// 	result := "C"
// 	// 	println(result)
// 	// }
// 	// println(result)

// 	//スコープの外と中で同じresultを共有したい場合は、:を外す
// 	ok, result := true, "A"
// 	fmt.Printf("memory address of result: %p\n", &result) //119行目と122行目のresultの番地が一緒
// 	if ok {
// 		result = "B"
// 		fmt.Printf("memory address of result: %p\n", &result)
// 		println(result)
// 	} else {
// 		result = "C"
// 		println(result)
// 	}
// 	println(result)
// }
