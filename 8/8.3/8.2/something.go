package something

import "fmt"

func makeSomething(n int) []string {
	var r = make([]string, n, n)
	for i := 0; i < n; i++ {
		r[i] = fmt.Sprintf("%5d 何か", i)
		// r = append(r, fmt.Sprintf("%05d 何か", i))
	}
	return r
}
