// This is NON concurrent code
package main

import (
	"fmt"
	//"log"
	"runtime"
	//"sync"
	"time"
)



func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(16)                                                        // Extra processors don't help w sequential tasks!!!!!!

	start := time.Now()															  // a..h is 8 intesive serail workloads, each of 10 billion loops on CPU
	counta()
	countb()
	countc()
	countd()
	counte()
	countf()
	countg()
	counth()

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}
func counta() {
	fmt.Println("AAAA is starting  ")
	for i := 1; i < 10_000_000_000; i ++ {										// 10 billion loops, simulating CPU extensive work
	}

	fmt.Println("AAAA is done  ")


}
func countb() {
	fmt.Println("BBBB is starting  ")
	for i := 1; i < 10_000_000_000; i++ {										// 10 billion loops, simulating CPU extensive work
	}

	fmt.Println("BBBB is done")


}
func countc() {
	fmt.Println("CCCC is starting  ")
	for i := 1; i < 10_000_000_000; i++ {										// 10 billion loops, simulating CPU extensive work
	}

	fmt.Println("CCCC is done    ")


}
func countd() {
	fmt.Println("DDDD is starting  ")
	for i := 1; i < 10_000_000_000; i++ {										// 10 billion loops, simulating CPU extensive work
	}

	fmt.Println("DDDD is done     ")


}
func counte() {
	fmt.Println("EEEE is starting  ")
	for i := 1; i < 10_000_000_000; i++ {										// 10 billion loops, simulating CPU extensive work
	}

	fmt.Println("EEEE is done   ")


}
func countf() {
	fmt.Println("FFFF is starting  ")
	for i := 1; i < 10_000_000_000; i++ {										// 10 billion loops, simulating CPU extensive work
	}

	fmt.Println("FFFF is done     ")


}
func countg() {
	fmt.Println("GGGG is starting  ")
	for i := 1; i < 10_000_000_000; i++ {										// 10 billion loops, simulating CPU extensive work
	}

	fmt.Println("GGGG is done     ")


}
func counth() {
	fmt.Println("HHHH is starting  ")
	for i := 1; i < 10_000_000_000; i++ {										// 10 billion loops, simulating CPU extensive work
	}

	fmt.Println("HHHH is done     ")


}

//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:37815 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/8-IO_Bound_vs_CPU_Bound/CPU_1-sequential
// DAP server listening at: 127.0.0.1:37815
// 2												<----- max proc on my PC == 2
// AAAA is starting  
// AAAA is done       ----
// BBBB is starting   ----  happen almost at the same time
// BBBB is done
// CCCC is starting  
// CCCC is done    
// DDDD is starting  
// DDDD is done     
// EEEE is starting  
// EEEE is done   
// FFFF is starting  
// FFFF is done     
// GGGG is starting  
// GGGG is done     ----
// HHHH is starting ----  happen almost at the smae time 
// HHHH is done     
// Processes took 1m33.785319346s
// Process 93535 has exited with status 0
// dlv dap (93491) exited with code: 0