package main

import (
	"fmt"
	"reflect"
)

func main() {
	basic()
	interfaceAny()
}

type Value int

func (v *Value) Add(n Value) {
	*v += n
}

func basic() {
	{
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

	fmt.Println("-----------------------")
	{
		// iota(列挙)
		// iotaはギリシャ文字ιで、「小さい量」を意味する
		const (
			Apple  = iota        // 0
			Orenge = iota + iota // 1 + 1 = 2
			Banana               // 2 + 2 = 4
			Grape  = iota        // 3
		)
		fmt.Println(Apple, Orenge, Banana, Grape)
	}

	fmt.Println("-----------------------")
	{
		type Fruit int
		type Animal int

		const (
			Orenge Fruit = iota //型を指定できる
			Apple
			Banana Animal = iota // 別の方を混同させるのはよくないけどできる
			Cherry
		)

		fmt.Println(Orenge, Apple)
		fmt.Println(reflect.TypeOf(Orenge), reflect.TypeOf(Apple), reflect.TypeOf(Banana), reflect.TypeOf(Cherry))
	}
	fmt.Println("-----------------------")
	{
		i := 2
		switch i {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2 or ")
			fallthrough // 次のcaseも実行する
		case 3, 4:
			fmt.Println("3")
		default:
			fmt.Println("other")
		}

	}
	//スライス
	{
		a := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			a = append(a, i)
		}
		a2 := make([]int, 0, len(a))
		for i := 0; i < len(a); i++ {
			if i%2 == 0 {
				//奇数は削除する
				a2 = append(a2, a[i])
			}
		}
		a = a2
		fmt.Println(a)
	}
	{
		a := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			a = append(a, i)
		}

		n := 50
		a = append(a[:n], a[n+1:]...) //n番目を削除
		fmt.Println(a)
	}
	{
		a := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			a = append(a, i)
		}
		n := 50
		fmt.Println("a[n+1:]", a[n+1:])
		fmt.Println("a[:n]", a[:n])
		a = append(a[:n+copy(a[n:], a[n+1:])])
		fmt.Println(a, len(a))

		// a := []int{1, 2, 3}
		// b := []int{4, 5}
		// fmt.Println(copy(a, b), a, b)
	}
	//文字列
	{
		name := "ami"
		fmt.Printf("%c", name[0])
		// fmt.Println(name, name[0])
		// name[0] = "A" // 文字列はイミュータブルなのでエラーになる
	}
	{
		s := "Hello"
		b := []byte(s)
		b[0] = 'h'
		s = string(b) //イミュータブルなので再代入が必要
	}
	{
		s := "こんにちわ世界"
		rs := []rune(s) //rune型でUnicodeのコードポイント列に変換できる
		rs[4] = 'は'
		s = string(rs)
	}
	//map
	{
		m := make(map[string]int, 1)
		m = map[string]int{
			"apple":  100,
			"banana": 200,
		}
		if v, ok := m["grape"]; !ok {
			fmt.Println(v, ok)
		}
		for k, v := range m {
			fmt.Printf("key:%v, value:%v\n", k, v)
		}
	}
	//型宣言
	{
		type MySgring string
		a := "hello"
		var ms MySgring = MySgring(a)
		fmt.Println(ms)
	}
	// 構造体
	{
		// 関数の引数などにstructを渡すと、都度コピーが行われるので、
		// コピーのオーバーヘッドをなくすのならポインタを使うといい
		type user struct {
			Name string
			age  int
		}

		showName := func(user *user) {
			fmt.Println(user.Name)
		}

		u := user{
			Name: "ami",
			age:  33,
		}
		showName(&u)
	}
	// ポインタ
	{
		type User struct {
			Name string
			Age  int
		}
		u := new(User) //ポインタを返す
		u.Name = "ami"
		u.Age = 33
	}
	{

		v := Value(1)
		v.Add(2)
		fmt.Println(v)
	}
}

func interfaceAny() {
	var v interface{}
	v = 1
	v = "こんにちは"
	n := v.(string)
	fmt.Println(n)

	//間違った方で型アサーションするとpanicが発生するので、以下のように確認する
	s, ok := v.(int)
	fmt.Println(s)
	if !ok {
		fmt.Println("vはintではない")
	}
}
