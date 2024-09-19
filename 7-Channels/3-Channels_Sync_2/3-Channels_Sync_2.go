package main

import (
	"fmt"
	"time"
)

func main() {

	var value = 97
	var result = 0
	goChan := make(chan int)
	mainChan := make(chan string)
	calculateSquare := func() { // This is a nested function, assigned to a variable
		time.Sleep(time.Second * 3)
		result = value * value
		goChan <- result
	}
	reportResult := func() { // This is a nested function, assigned to a variable
		fmt.Println(value, "squared is", <-goChan)                                 // <------------blocks on <-mainChan, on unblock prints Rcv's side of  <-goChan
		// blocks until it can read something from goChan - printed
		mainChan <- "You can quit now.  I'm done." // This is just for clarity.
		fmt.Println(<-mainChan) //discarded is string "You can quit now. I'm done."
	}

	go calculateSquare()
	go reportResult()
	<-mainChan // blocks until it can read something from mainChan - discarded


}
//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:35803 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/3-Channels_Sync_2
// DAP server listening at: 127.0.0.1:35803
// 97 squared is 9409
// Process 73912 has exited with status 0
// dlv dap (73869) exited with code: 0