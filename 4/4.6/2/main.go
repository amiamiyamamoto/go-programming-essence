package main

import (
	"bp/server"
	"log"
	"os"
	"time"
)

func main() {
	// f, err := os.Create("server.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// logger := log.New(f, "", log.LstdFlags)
	// svr := server.NewBuilder("localhost", 8888).
	// 	Timeout(time.Minute).
	// 	Logger(logger).
	// 	Build()

	builder := server.NewBuilder("localhost", 8888)
	configureServer(builder)
	svr := builder.Build()
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}

func configureServer(build *server.ServerParam) {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	to := time.Duration(time.Minute)

	logger := log.New(f, "", log.LstdFlags)
	build.Logger(logger).Timeout(to)
}
