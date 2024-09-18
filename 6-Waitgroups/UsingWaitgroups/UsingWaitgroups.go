package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // wg := 0

func main() {

	wg.Add(2) // wg := 2
	start := time.Now()
	go doSomething()     // wg -= 1
	go doSomethingElse() // wg -= 1
	wg.Wait()            // wg := 2 (while waiting), 1 (after first decrement), 0 (after 2nd decrement)

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("\nI've done something")
	wg.Done() // explicit derement of wg counter
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
	wg.Done() // explicit derement of wg counter
}

//OUT:
//proper synchbronization with Waigroup!!!!
//
//Starting: /go/bin/dlv dap --listen=127.0.0.1:39865 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/6-Waitgroups/UsingWaitgroups
// DAP server listening at: 127.0.0.1:39865
//
// I've done something
// I've done something else
//
//
// I guess I'm done
// Processes took 2.000233904s						<----both go routines execute in parallel, since eachis modeled to execute ~ 2sec => both execute in ~2secs+
// Process 15753 has exited with status 0
// dlv dap (15708) exited with code: 0
