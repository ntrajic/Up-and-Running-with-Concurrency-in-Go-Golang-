package main

import (
	"fmt"
	"time"

)


func main(){


	start := time.Now()
	go doSomething()
	go doSomethingElse()





	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething(){
	time.Sleep(time.Second*2)
	fmt.Println("\nI've done something")
}


func doSomethingElse(){
	time.Sleep(time.Second*2)
	fmt.Println("I've done something else")
}

// OUT:
// no synchornization!!!!!
//
//Starting: /go/bin/dlv dap --listen=127.0.0.1:35109 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/4-First_Goroutine/BadAttempt1UsingGoroutines
// DAP server listening at: 127.0.0.1:35109
//
//
// I guess I'm done
// Processes took 19.847µs
// Process 10653 has exited with status 0
// dlv dap (10615) exited with code: 0