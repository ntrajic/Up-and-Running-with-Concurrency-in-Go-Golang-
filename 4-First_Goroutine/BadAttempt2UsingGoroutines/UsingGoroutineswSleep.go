package main

import (
	"fmt"
	"time"

)


func main(){

	start := time.Now()
	go doSomething()
	go doSomethingElse()

	time.Sleep(time.Second*5)  						// Don't do this in production.  VERY inefficient and unpredictable.!!!!!

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

//OUT:
// Don't do this in production.  VERY inefficient and unpredictable.!!!!!
//
//Starting: /go/bin/dlv dap --listen=127.0.0.1:44811 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/4-First_Goroutine/BadAttempt2UsingGoroutines
// DAP server listening at: 127.0.0.1:44811
//
// I've done something
// I've done something else
//
//
// I guess I'm done
// Processes took 5.002605911s					// both go routines execute about 2sec => 2+2 < 5 (sleep), so both will execute while main is sleeping for 5 se3conds
// Process 12454 has exited with status 0
// dlv dap (12410) exited with code: 0