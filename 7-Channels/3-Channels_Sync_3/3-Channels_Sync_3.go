package main

import (
	"fmt"
	"sync"
	"time"
)



func main() {

	var wg sync.WaitGroup // wg := 0
	wg.Add(1)			  // wg += 1, wg is 1
	go func() {  // This turns our entire process into One Large Goroutine.
		var value = 97
		var result = 0
		goChan := make(chan int)
		mainChan := make(chan string)
		calculateSquare := func() {
			time.Sleep(time.Second * 3)
			result = value * value
			goChan <- result
		}
		reportResult := func() {
			fmt.Println(value, "squared is", <-goChan)
			// blocks until it can read something from goChan - printed
			mainChan <- "You can quit now.  I'm done." // This is just for clarity.
		}

		go calculateSquare()
		go reportResult()
		<-mainChan // blocks until it can read something from mainChan - discarded
		wg.Done()  // wg -= 1 , wg is 0        <------ last line in Large Goroutine decrements wg to signal it's done - read the result from goChan
	}()

	wg.Wait()	   // blocks until whole One Large Goroutine is done (wg is 0), wg is blocked/waits while wg == 1
}
