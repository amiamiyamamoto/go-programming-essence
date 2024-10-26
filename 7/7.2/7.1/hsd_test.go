package hsd

import "testing"

func TestStiringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 1
	if got != want {
		t.Fatalf("expected %v, got %v:", want, got)
	}
}
