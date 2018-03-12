package main

import (
	_ "net/http/pprof"
	"jobShop/test"
	"flag"
	"os"
	"log"
	"runtime/pprof"
)

func main() {
	flag.Parse()

	var cpuprofile = flag.String("cpuprofile", "build/cpu2.prof", "write cpu profile to `file`")
	//var memprofile = flag.String("memprofile", "mem.prof", "write memory profile to `file`")

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

	test.TestAll()
	return


	//jobs := base.Jobs{
	//	base.Job{
	//		base.Task{0, 3},
	//		base.Task{1, 2},
	//		base.Task{2, 2},
	//	},
	//
	//	base.Job{
	//		base.Task{Machine: 0, Time: 2},
	//		base.Task{Machine: 2, Time: 1},
	//		base.Task{Machine: 1, Time: 4},
	//	},
	//
	//	base.Job{
	//		base.Task{1, 4},
	//		base.Task{2, 3},
	//	},
	//}
	//
	//newState := state.NewState(jobs)
	//solution := simpleGreed.Resolver{State: newState}.FindSolution()
	//resultState := solution.(simpleGreed.Resolver).State
	//
	//solver := tabuSearch.NewSolver(resultState)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(solver.BestMakespan())
	//	solver.Next()
	//}

	//if *memprofile != "" {
	//	f, err := os.Create(*memprofile)
	//	if err != nil {
	//		log.Fatal("could not create memory profile: ", err)
	//	}
	//	runtime.GC() // get up-to-date statistics
	//	if err := pprof.WriteHeapProfile(f); err != nil {
	//		log.Fatal("could not write memory profile: ", err)
	//	}
	//	f.Close()
	//}
}
