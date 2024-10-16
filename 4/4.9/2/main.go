package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}
type AttrEx struct {
	Name string
}

type Teacher struct {
	Attr
	AttrEx
	Subject string
}

func main() {
	teacher := Teacher{
		Attr: Attr{
			Name: "John Schwartz",
			Age:  43,
		},
		AttrEx: AttrEx{
			Name: "JS",
		},
		Subject: "Math",
	}
	// fmt.Println(teacher.Name)//フィールド名が重複しているので直接は呼べない
	fmt.Println(teacher.Attr.Name)

}
