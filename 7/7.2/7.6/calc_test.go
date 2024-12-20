package calc

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		lhs  int
		rhs  int
		want int
	}{
		{name: "test1", lhs: 0, rhs: 1, want: 1},
		{name: "test2", lhs: 1, rhs: -1, want: 0},
		{name: "test3", lhs: 2, rhs: 1, want: 3},
	}

	for _, test := range tests {
		got := Add(test.lhs, test.rhs)
		if got != test.want {
			t.Errorf("%v: want %v,but %v", test.name, test.want, got)
		}
	}
}
