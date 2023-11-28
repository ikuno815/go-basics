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

	//配列は動的に要素追加ができないのでsliceを使う
	//要素数を空にすることでsliceを定義していくことができる
	var s1 []int  //s1がnilとして認識される
	s2 := []int{} //nilとして認識されない
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))

	//varと:=の違いは、varはnilとして認識されるが:=だとnilとして認識されないこと
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //false

	//要素の追加はappendを使うことができる
	//s1に対して、123という値を追加する
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	//s1にs3という別で定義したsliceを追加する
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	//s3の末尾に...をつけるとこのsliceを可変長の引数の形でappendに変換して渡すことができる
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	//makeを使ったsliceの定義方法
	//makeでは、指定したcapacityの値に応じてメモリ領域を確保してくれる
	s4 := make([]int, 0, 2) //要素数が0でcapacityが2のint型のsliceをmakeで作っている
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	s4 = append(s4, 1, 2, 3, 4)
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))

	//0 4つで初期化される
	s5 := make([]int, 4, 6)
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	//1と2のindexを取得したい場合は、1から3にする（上限のインデックスはプラス1で表現されるため）
	s6 := s5[1:3]
	s6[1] = 10
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))

	//sliceを切り取って使用した場合はメモリを共有する仕様になっている
	//s6の値を変えるとs5の方も変わる
	s6 = append(s6, 2)
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))

	//メモリを共有させたくない場合はcopyを使う
	sc6 := make([]int, len(s5[1:3])) //copy元の長さを指定する
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6 dst copy before: %v %v %v\n", sc6, len(sc6), cap(sc6))

	copy(sc6, s5[1:3]) //左にcopy先のslice、右にcopy元のsliceを指定する
	fmt.Printf("sc6 dst of copy after: %v %v %v\n", sc6, len(sc6), cap(sc6))

	sc6[1] = 12
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5)) //s5は変更されない
	fmt.Printf("sc6: %v %v %v\n", sc6, len(sc6), cap(sc6))

	s5 = make([]int, 4, 6) //0を4つで初期化
	fs6 := s5[1:3:3]       //メモリの共有を許可する最大のindexを3つ目の値で指定することができる
	//↑指定した値-1のindex（今回は３なので２）までメモリを共有するということ
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5)) //8はindex3なのでs5には反映されない
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))

	s5[3] = 9
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5)) //s5のindex3には9が反映される
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))

	//map: keyと値をテーブルとして管理できるもの
	//定義方法：mapというkeywordを使って、まず[]のところでkeyの型を指定する
	//右側には値を型を指定する
	//ここではkeyがstring型で、値がintの型のmapを定義している
	var m1 map[string]int
	//:=でも定義できるが、その場合は空の{}で空の値を代入する必要がある
	m2 := map[string]int{}
	fmt.Printf("%v %v \n", m1, m1 == nil) //map[] nil
	fmt.Printf("%v %v \n", m2, m2 == nil) //map[] nilではない

	//m2のmapに値を設定してみる
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"])

	//deleteを使って要素を削除してみる
	delete(m2, "A") //Aのkeywordに紐づいた要素を削除する
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"])

	//存在する0("C"の値)と存在しない0の識別方法
	v, ok := m2["A"]             //返り値の２番目にok(true,false)が返ってくる
	fmt.Printf("%v %v\n", v, ok) //0 false
	v, ok = m2["C"]
	fmt.Printf("%v %v\n", v, ok) //0 true

	//mapの要素をfor文で取り出す方法
	//rangeを使うことができて、mapの要素を1個1個取り出してfor文を回してくれる
	//mapの注意点：内部でhashmapを使っているのでこのforで取り出す順番が必ず同じ順番になるとは限らない！
	for k, v := range m2 {
		fmt.Printf("%v %v\n", k, v)
	}
}
