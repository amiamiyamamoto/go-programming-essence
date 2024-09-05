package main

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func main() {
	user := NewUser("John", 30)
	println(user.Name, user.Age)
}
