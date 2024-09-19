//https://www.youtube.com/watch?v=LvgVSSpwND8&t=563s
//From Youtube  "Concurrency in Go" tutorial by Jake Wright

package main

import (
	"fmt"
)

func main() {

	c := make(chan string, 3)  // channel doesn't block until full ("buffered" channel = NON-BLOCKING channel), capacity of channels = 3   
	c <- "Hello "
	c <- "Earth "
	c <- "from Mars"
	//c <- "from Venus"		   // NOTE: if 4th string "from Venus" is sent, channel c goes over its capacity = 3, than channel blocks, and DEADLOCK happens

	msg := <-c  			   // DECLARATION  w/ :=
	fmt.Print(msg)			   // print 1st msg ("Hello")

	msg = <-c 				   // Notice we used = NOT := because msg is already declared
	fmt.Print(msg)             // print 2nd msg ("Earth")

	msg = <-c  				   // Notice we used = NOT := because msg is already declared
	fmt.Println(msg)           // print 3rd msg ("from Mars")

}


//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:40319 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/4-BuffferedChannels
// DAP server listening at: 127.0.0.1:40319
// Hello Earth from Mars							<------- 3 times sending different strings to channel c
// Process 52986 has exited with status 0
// dlv dap (52940) exited with code: 0


//when 4th message "from Venus" is sent, i.e. line 16 is uncommented, OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:44551 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/4-BuffferedChannels
// DAP server listening at: 127.0.0.1:44551
// fatal error: all goroutines are asleep - deadlock!   <----- this s b/c non-blocking channel c has capacity=3, and ve went over it, all coros go to sleeop <=> deadlock
//
// goroutine 1 [chan send]:
// main.main()
// 	/workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/4-BuffferedChannels/4-bufferedChannels.go:16 +0x6f
// Process 54112 has exited with status 2
// dlv dap (54069) exited with code: 0