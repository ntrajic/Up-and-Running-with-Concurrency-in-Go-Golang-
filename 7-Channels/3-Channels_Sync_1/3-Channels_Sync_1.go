package main

import (
	"fmt"
	"time"
)

var result = 0
var value = 97

func main() {
	goChan := make(chan int)
	mainChan := make(chan string)
	go calculateSquare(value, goChan)
	go reportResult(goChan, mainChan)
	<-mainChan // blocks until it can read something from mainChan - discarded
}
func calculateSquare(value int, goChan chan int) {
	fmt.Println("Calculating for 3 seconds...")
	time.Sleep(time.Second * 3)
	result = value * value
	goChan <- result

}
func reportResult(goChan chan int, mainChan chan string) {
	time.Sleep(time.Second * 1)
	fmt.Println("The result of", value, "squared", "is", <-goChan)
	// blocks until it can read something from goChan - printed
	mainChan <- "You can quit now.  I'm done." // This is just for clarity.
}

//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:33491 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/3-Channels_Sync_1
// DAP server listening at: 127.0.0.1:33491
// Calculating for 3 seconds...
// The result of 97 squared is 9409
// Process 73214 has exited with status 0
// dlv dap (73169) exited with code: 0