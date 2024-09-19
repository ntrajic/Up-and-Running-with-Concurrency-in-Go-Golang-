// Uses a select / case statement to monitor 2 channels and print whichever is active

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	c1 := make(chan string)				// 3 blocking channels, read in a "switch" select statment, in an infinite loop
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			c1 <- "Sending every 1 second"

		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 4)
			c2 <- "Sending every 4 sec"

		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 10)
			c3 <- "We're done"
		}
	}()

	for { // infinite for loop  This is the operator - listening for activity on all channels.
		// This is a clever way to get around the blocking nature when you try to read a channel
		select {
		case msg := <-c1:
			fmt.Println(msg)									// reading c1 every 1 sec
		case msg := <-c2:
			fmt.Println(msg + " Something cool happened")		// reading c2 every 4 sec
		case msg := <-c3:				
			fmt.Println(msg)									// reading c3 after 10 secs, then exit.					
			os.Exit(0)

		}

	}
}

//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:44415 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/5-MultipleChannels_Select
// DAP server listening at: 127.0.0.1:44415
// 3  Sending every 1 second 	                   // (channel c1)
//    Sending every 4 sec Something cool happened  // (4th second)                  
// 4 Sending every 1 second                        // (4,5,6,7,8th second, channel c2)
//   Sending every 4 sec Something cool happened   // (8th second, channel c1)
// 2 Sending every 1 second                        // (9th,10h, channel c1)
// We're done                                      // (channel c3)
// Process 61987 has exited with status 0
// dlv dap (61942) exited with code: 0