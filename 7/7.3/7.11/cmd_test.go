package cmd_test

import "testing"

func TestCreateProfile(t *testing.T) {
	t.Setenv("DATABASE_URL", "test")
	err := doSomething()
	if err != nil {
		t.Fatalf("cannot do something: %v", err)
	}

}

func doSomething() error {
	return nil
}
