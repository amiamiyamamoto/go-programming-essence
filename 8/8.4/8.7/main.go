package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func hevyFunc(wf *sync.WaitGroup) {
	defer wf.Done()
	s := make([]string, 3)
	for i := 0; i < 1000000; i++ {
		s = append(s, "magicval pandas")
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile) //(1)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile() //(2)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go hevyFunc(&wg)
	wg.Wait()

	if *memprofile != "" {
		f, err := os.Create(*memprofile) //(3)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() //(4)
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
