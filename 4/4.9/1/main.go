package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}
type Teacher struct {
	Attr
	Subject string
}
type Student struct {
	Attr
	Score int
}

func main() {
	teacher := Teacher{
		Attr: Attr{
			Name: "Taro",
			Age:  43,
		},
		Subject: "Math",
	}
	student := Student{
		Attr: Attr{
			Name: "Hanako",
			Age:  15,
		},
		Score: 90,
	}
	//合成したstructのフィールドもちょくでアクセスできる
	fmt.Println(teacher.Name)
	fmt.Println(student.Name)
}
