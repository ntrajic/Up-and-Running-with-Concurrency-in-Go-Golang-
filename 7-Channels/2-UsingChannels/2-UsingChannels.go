package main

import (
	"fmt"
	"time"

)
var ch = make(chan string)							// blocking

func main(){


	start := time.Now()
	go doSomething()
	go doSomethingElse()


	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)							// should be ~2+ secs, b/c both coroutines executed in parallel!!!
	fmt.Printf("Processes took %s", elapsed)
}

// both coroutines execute ~2+ secs
func doSomething(){
	time.Sleep(time.Second*2)
	fmt.Println("\nI've done something")
	ch <- "doSomething finished"
}

func doSomethingElse(){
	time.Sleep(time.Second*2)
	fmt.Println("I've done something else")
	ch <- "doSomethingElse finished"
}

//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:43151 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/2-UsingChannels
// DAP server listening at: 127.0.0.1:43151
// I've done something else
// doSomethingElse finished
//
// I've done something
// doSomething finished
// I guess I'm done
// Processes took 2.001155397s							<------- ~2+ secs, b/c bothn coros executed in parallel
// Process 47199 has exited with status 0
// dlv dap (47155) exited with code: 0