package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Before")
	ret := m.Run()
	log.Println("After")

	fmt.Println(ret)
	os.Exit(ret)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
