package main

import "fmt"

//sliceとmapについて
//配列

func main() {
	//明示的な代入がない場合はゼロ値で初期化される
	var a1 [3]int
	//明示的な代入がある場合はcurly bracketで表現する
	var a2 = [3]int{10, 20, 30}
	// :=で定義する場合はcurly bracketでの定義が必ず必要になる
	// [...]スリードットをつけると、代入された値をカウントして自動的に配列の要素数を特定する
	a3 := [...]int{10, 20}
	fmt.Printf("%v %v %v\n", a1, a2, a3)

	//配列やsliceに関してはlengthとcapacityというメソッドが準備されている
	//capacityは対象の配列やsliceが利用できるメモリの量のこと
	fmt.Printf("%v %v\n", len(a3), cap(a3))

	//配列の方を表示
	fmt.Printf("%T %T\n", a2, a3)
}
