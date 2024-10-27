package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		log.Fatal("failed to setup:", err)
	}

	ret := m.Run()

	if err := teardown(); err != nil {
		log.Fatal("failed to teardown:", err)
	}
	os.Exit(ret)
}

func setup() error {
	log.Println("setup")
	return nil
}

func teardown() error {
	log.Println("teardown")
	return nil
}
