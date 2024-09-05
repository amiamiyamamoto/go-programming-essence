package main

import (
	"fmt"
	"reflect"
)

// どんな型でも受け取れる関数
// typeで定義した型を扱うことはできない
func PrintDetail(v interface{}) {
	switch t := v.(type) {
	case int, int32, int64:
		fmt.Println("int/int32/int64 型:", t)
	case string:
		fmt.Println("string 型:", t)
	default:
		fmt.Println("知らない型")
	}
}

// type定義された型を扱うにはreflectパッケージを使う
func PrintDetailOrgtype(v interface{}) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int/iint32/int64 型:", v)
	case reflect.String:
		fmt.Println("string 型:", v)
	default:
		fmt.Println("知らない型")
	}
}

func main() {
	type V int
	var v V = 100
	PrintDetail(v)
	PrintDetailOrgtype(v)
}
