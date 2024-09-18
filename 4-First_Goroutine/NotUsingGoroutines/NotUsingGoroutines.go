package main

import (
	"fmt"
	"time"

)


func main(){

	start := time.Now()
	go doSomething()
	go doSomethingElse()

	time.Sleep(time.Second*5)

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
//Starting: /go/bin/dlv dap --listen=127.0.0.1:37247 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/4-First_Goroutine/NotUsingGoroutines
// DAP server listening at: 127.0.0.1:37247
// I've done something else
//
// I've done something
//
//
// I guess I'm done
// Processes took 5.001552527s						<---- it takes 5+ secs to execute non-synchronized go routines!!!
// Process 4687 has exited with status 0
// dlv dap (4641) exited with code: 0