package main

import (
	"fmt"
	"reflect"
)

func main() {

	const n = 1
	//constは使われる場所でそれぞれ型が決まるためコンパイルエラーにならない
	x := 1 + n
	y := 1.2 + n

	z := 9223372036854775807
	a := 1. //これでフロートになる
	//それぞれの型を出力
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y), reflect.TypeOf(z), a)

	//intサイズを調べる
	//intサイズは実行環境のCPUが32ビットか64ビットかで変わる
	const intSize = 32 << (^uint(0) >> 63)
	fmt.Println(intSize)
}
