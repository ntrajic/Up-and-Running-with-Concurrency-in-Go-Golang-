package main

import (
	"fmt"
	"sync"
	)

var (
	wg sync.WaitGroup
	mutex = sync.Mutex{}
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
		mutex.Lock()
		widgetInventory -= 100
		mutex.Unlock()
	}

	wg.Done()
}

func newPurchases() {  // 1000000 widgets purchased
	for i := 0; i < 300000; i++ {
		mutex.Lock()
		widgetInventory+= 100
		mutex.Unlock()
	}
	wg.Done()
}
//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:42951 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/9-Race_Conditions/2-raceConditionMutex
// DAP server listening at: 127.0.0.1:42951
// Starting inventory count =  1000
// Ending inventory count =  1000
// Process 135180 has exited with status 0
// dlv dap (135136) exited with code: 0