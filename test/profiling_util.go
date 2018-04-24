package test

import (
	_ "net/http/pprof"
	"flag"
	"os"
	"log"
	"runtime/pprof"
	"runtime"
)

func Profile(it func()) {
	flag.Parse()

	var cpuprofile = flag.String("cpuprofile", "build/cpu2.prof", "write cpu profile to `file`")
	var memprofile = flag.String("memprofile", "build/mem.mprof", "write memory profile to `file`")

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	it()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
