package main

import (
	"log"
	"testing"
)

func TestA(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	log.Println("TestA running")
}
