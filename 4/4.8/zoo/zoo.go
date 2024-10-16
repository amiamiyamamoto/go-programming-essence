package zoo

import (
	"internaltest/zoo/internal/bored"
)

func DoZoo() {
	println("I'm in the zoo")
	println(bored.DoBored()) //zooからは呼べる
}
