package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10000; i++ {  // This allows us to create MANY (10000) large goroutines, it will all finish in 3 secs in parallel!!!
		wg.Add(1)
		go func(j int) {

			var result = 0
			goChan := make(chan int)
			mainChan := make(chan string)
			calculateSquare := func() {
				time.Sleep(time.Second * 3)  // Deliberate time delay
				result = j * j
				goChan <- result
			}
			reportResult := func() {
				fmt.Println(j, "squared is", <-goChan)
				// blocks until it can read something from goChan - printed
				mainChan <- "You can quit now.  I'm done." // This is just for clarity.
			}

			go calculateSquare()
			go reportResult()
			<-mainChan // blocks until it can read something from mainChan - discarded
			wg.Done()
		}(i)
	}
	wg.Wait()
}
//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:40697 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/7-Channels/3-Channels_Sync_4
//DAP server listening at: 127.0.0.1:40697
// 34 squared is 1156
// 17 squared is 289
// 3 squared is 9
// 19 squared is 361
// 4 squared is 16
// 20 squared is 400
// 6 squared is 36
// 22 squared is 484
// 7 squared is 49
// 23 squared is 529
// 8 squared is 64
// ...                              <----- 10000 results!!!! in 3 secs
// 8861 squared is 78517321
// 6742 squared is 45454564
// 9467 squared is 89624089
// 9629 squared is 92717641
// 6736 squared is 45373696
// 6740 squared is 45427600
// 5701 squared is 32501401
// 6746 squared is 45508516
// Process 84206 has exited with status 0
// dlv dap (84163) exited with code: 0