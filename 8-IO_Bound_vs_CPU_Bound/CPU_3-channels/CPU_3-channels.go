// This is concurrency / goroutines implemented with CHANNELS
package main

import (
	"fmt"
	"runtime"

	"time"
)



func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(16)   //  Extra processors help up to number of goroutines w CPU-bound tasks
	c := make(chan string)
	startTime := time.Now()
	go counta(c)
	go countb(c)
	go countc(c)
	go countd(c)
	go counte(c)
	go countf(c)
	go countg(c)
	go counth(c)



	for i := 0; i < 8; i++ {
		fmt.Println(<-c)
		}

	elapsed := time.Since(startTime)
	fmt.Printf("Processes took %s", elapsed)
}
func counta(c chan string) {
	fmt.Println("AAAA is starting  ")
	for I := 1; I < 10_000_000_000; I ++ {
	}


	c <- "AAAA is done"

}
func countb(c chan string) {
	fmt.Println("BBBB is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}


	c <- "BBBB is done"

}
func countc(c chan string) {
	fmt.Println("CCCC is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "CCCC is done"


}
func countd(c chan string) {
	fmt.Println("DDDD is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}


	c <- "DDDD is done"

}
func counte(c chan string) {
	fmt.Println("EEEE is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "EEEE is done"


}
func countf(c chan string) {
	fmt.Println("FFFF is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}


	c <- "FFFF is done"

}
func countg(c chan string) {
	fmt.Println("GGGG is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "GGGG is done"


}
func counth(c chan string) {
	fmt.Println("HHHH is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "HHHH is done"


}

//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:33973 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/8-IO_Bound_vs_CPU_Bound/CPU_3-channels
// DAP server listening at: 127.0.0.1:33973
// 2
// BBBB is starting  
// AAAA is starting  
// DDDD is starting     
// FFFF is starting     
// HHHH is starting     
// GGGG is starting     
// CCCC is starting     
// EEEE is starting     
// HHHH is done
// BBBB is done
// FFFF is done
// DDDD is done
// CCCC is done
// GGGG is done
// EEEE is done
// AAAA is done
// Processes took 1m33.979334774s                       <----- should have shown speed up with coros and 1 channel, authors machine crunches in 3 secs w/ coros and channel
// Process 101191 has exited with status 0
// dlv dap (101147) exited with code: 0