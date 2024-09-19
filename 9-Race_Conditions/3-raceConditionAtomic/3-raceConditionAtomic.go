package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
var (
	wg sync.WaitGroup

	widgetInventory int32= 1000  //Package-level variable will avoid all the pointers
)


func main() {
	fmt.Println("Starting inventory count = ", widgetInventory)
	wg.Add(2)
	go makeSales()
	go newPurchases()
	wg.Wait()
	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales() {  // 1000000 widgets sold
	for i := 0; i < 300000; i++ {

		atomic.AddInt32(&widgetInventory,-100)			// atomic decrement

	}

	wg.Done()
}

func newPurchases() {  // 1000000 widgets purchased
	for i := 0; i < 300000; i++ {

		atomic.AddInt32(&widgetInventory,100)			// atomic increment

	}
	wg.Done()
}
// OUT: 
//Starting: /go/bin/dlv dap --listen=127.0.0.1:38433 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/9-Race_Conditions/3-raceConditionAtomic
// DAP server listening at: 127.0.0.1:38433
// Starting inventory count =  1000
// Ending inventory count =  1000
// Process 135563 has exited with status 0
//dlv dap (135520) exited with code: 0