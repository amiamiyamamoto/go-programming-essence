package dosomething_test

import (
	"fmt"
	"testing"
)

func FuzzDoSomething(f *testing.F) {
	f.Add("test&&&")
	f.Fuzz(func(f *testing.T, s string) {
		// fmt.Println(s)
		doSomething(s)
	})
}

func doSomething(s string) {
	fmt.Println(s)
	// if len(s) > 1 {
	// 	panic("error")
	// }
}

// func FuzzDoSomething2(f *testing.F) {
// 	f.Add(3, "test&&&")
// 	f.Fuzz(func(f *testing.T, i int, s string) {
// 		// fmt.Println(s)
// 		doSomething2(i, s)
// 	})
// }

// func doSomething2(i int, s string) {
// 	fmt.Println(s + string(i))
// }
