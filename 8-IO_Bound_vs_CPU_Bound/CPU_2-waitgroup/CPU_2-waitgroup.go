// This is concurrency / goroutines implemented with a WAITGROUP
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(32)     //  Extra processors help up to number of goroutines w CPU-bound tasks
	wg.Add(8)  // create 8 contractor badges for our workers
	startTime := time.Now()
	go counta() // pass out a badge to each worker
	go countb()
	go countc()
	go countd()
	go counte()
	go countf()
	go countg()
	go counth()





	wg.Wait()  // Do not end the program until all badges have been returned, ie all go routines have reported that they are done.
	elapsed := time.Since(startTime)
	fmt.Printf("Processes took %s", elapsed)
}
func counta() {
	fmt.Println("AAAA is starting  ")
	for I := 1; I < 10_000_000_000; I ++ {
	}

	fmt.Println("AAAA is done  ")
	wg.Done()  // Turn in my badge - I'm done

}
func countb() {
	fmt.Println("BBBB is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("BBBB is done")
	wg.Done()

}
func countc() {
	fmt.Println("CCCC is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("CCCC is done    ")
	wg.Done()

}
func countd() {
	fmt.Println("DDDD is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("DDDD is done     ")
	wg.Done()

}
func counte() {
	fmt.Println("EEEE is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("EEEE is done   ")
	wg.Done()

}
func countf() {
	fmt.Println("FFFF is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("FFFF is done     ")
	wg.Done()

}
func countg() {
	fmt.Println("GGGG is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("GGGG is done     ")
	wg.Done()

}
func counth() {
	fmt.Println("HHHH is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("HHHH is done     ")
	wg.Done()

}

// OUT:
// Starting: /go/bin/dlv dap --listen=127.0.0.1:34905 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/8-IO_Bound_vs_CPU_Bound/CPU_2-waitgroup
// DAP server listening at: 127.0.0.1:34905
// 2
// HHHH is starting  
// CCCC is starting  
// EEEE is starting  
// DDDD is starting  
// FFFF is starting  
// AAAA is starting  
// BBBB is starting  
// GGGG is starting  
// GGGG is done     
// CCCC is done    
// AAAA is done  
// EEEE is done   
// DDDD is done     
// BBBB is done
// HHHH is done     
// FFFF is done     
// Processes took 1m19.947554051s                            <----- author's machine crunches in 3 sec with goroutines and waitgroup, much faster than in sequential case!!!
// Process 96825 has exited with status 0
// dlv dap (96783) exited with code: 0