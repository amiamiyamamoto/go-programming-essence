package hsd

import "testing"

func TestStiringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 2 //1文字違いなのでこのテストは失敗する
	if got != want {
		t.Fatalf("expected %v, got %v:", want, got)
	}
}
