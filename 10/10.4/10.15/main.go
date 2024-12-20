package main

import (
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger(logs.AdapterMultiFile, `{
		"filename": "./foo.log",
		"daily": true,
		"maxlines": 10000,
		"rotate": true,
	}`)

	for i := 0; i < 100; i++ {
		log.Info("Hello, World!")
	}
}
