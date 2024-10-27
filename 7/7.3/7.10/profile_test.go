package prilife_test

import (
	"path/filepath"
	"testing"
)

func TestCreatePrifile(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "profile.json")
	got, err := CreateProfile(filename)
	if err != nil {
		t.Fatal(err)
	}
	want := true
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}
}

func CreateProfile(filename string) (bool, error) {
	return true, nil
}
