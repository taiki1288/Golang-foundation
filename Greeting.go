package main

import (
	"fmt"
)

// 変数は数字から始めることができない！
// アンダースコア以外に記号を使うことはできない！

func main() {
	num := 1
	num01 := 2
	num_01 := 3
	num$01 := 4

	fmt.Println(num$01)
}