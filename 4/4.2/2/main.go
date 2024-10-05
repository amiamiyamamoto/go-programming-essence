package main

import "fmt"

type F struct {
	Name string
	Age  int
}

func (f *F) String() string {
	return fmt.Sprintf("NAME=%q, AGE=%d", f.Name, f.Age)
}

func main() {
	f := &F{
		Name: "John",
		Age:  20,
	}

	fmt.Printf("%v\n", f)
	// Stringerインターフェースを実装しているので
	// NAME="John", AGE=20
	// と出力される

	a := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  20,
	}
	fmt.Printf("%v\n", a)
	// {John 20}
	// と出力される

}
