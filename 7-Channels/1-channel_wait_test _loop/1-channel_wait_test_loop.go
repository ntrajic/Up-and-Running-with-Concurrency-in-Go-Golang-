package main

import (
	"fmt"
	"time"
)



func main() {


	ch := make(chan string)	   // blocking channel

	//ch <- "hello from main"  //This won't work since it blocks, it's in the main coroutine

	// main coroutine reads on the R: rcv's end w/ fmt.Print(<-ch)    <----channel------  T: sendMe(ch) coroutine writes arbitrariry string in its body to channel ch
	go sendMe(ch)			   // put "SendMe is done" as string argument of sendMe(), arg:ch=channel,  "ch chan<- string",  body of sendMe() passes the string: ch <- "SendMe is done"
	
	
	for i:=1; i<2; i++{  // This for loop just reads the channel as messages come in, it is not neccessary, since whole string "SendMe is done" is wrtiten to Rcv in i:=1
		fmt.Println(<-ch)		   // print what is on the left side of arrow!!!, what is received
	}


}
func sendMe(ch chan<- string) {

	time.Sleep(time.Second*2)
	ch <- "SendMe is done"
}


// OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:32963 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/1-channel_wait_test
// DAP server listening at: 127.0.0.1:32963
// SendMe is done	                      <-------- this string is received, and printed on line 22, for i=1
// Process 31971 has exited with status 0
// dlv dap (31908) exited with code: 0
