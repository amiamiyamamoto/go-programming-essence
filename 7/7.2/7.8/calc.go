package calc

import "time"

func Add(a, b int) int {
	//すごく重たい足し算
	result := a + b
	time.Sleep(3 * time.Second)
	return result
}
