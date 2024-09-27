package main

var _ I = (*foo)(nil)

// (*foo)(nil)はfoo型のポインタ型のnilを表している
// interface I の型を持つ変数_にfoo型のポインタ型のnilを代入することで、foo型がinterface Iを実装しているかをコンパイラにチェックさせる

type I interface {
	doSomething()
}

type foo struct{}

func (f *foo) doSomething() {}
