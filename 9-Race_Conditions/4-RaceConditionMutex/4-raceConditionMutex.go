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
	wg.Wait()								//<<<<
	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales() {  // 1000000 widgets sold
	for i := 0; i < 3000; i++ {
		mutex.Lock()						//<<<
		widgetInventory -= 100
		fmt.Println(widgetInventory)
		mutex.Unlock()						//<<<<
	}

	wg.Done()
}

func newPurchases() {  // 1000000 widgets purchased
	for i := 0; i < 3000; i++ {
		mutex.Lock()						//<<<<
		widgetInventory+= 100
		fmt.Println(widgetInventory)
		mutex.Unlock()						//<<<<
	}
	wg.Done()
}
//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:36173 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/9-Race_Conditions/4-RaceConditionMutex
// DAP server listening at: 127.0.0.1:36173
// Starting inventory count =  1000
// 1100
// 1200
// 1300
// 1400
// 1500
// 1600
// 1700
// ...
// -5000
// -4900
// -4800
// -4700
// -4600
// -4500
// -4400
// -4300
// -4200
// -4100
// -4000
// -3900
// -3800
// -3700
// -3600
// -3500
// -3400
// -3300
// -3200
// -3100
// -3000
// -2900
// -2800
// -2700
// -2600
// -2500
// -2400
// -2300
// -2200
// -2100
// -2000
// -1900
// -1800
// -1700
// -1600
// -1500
// -1400
// -1300
// -1200
// -1100
// -1000
// -900
// -800
// -700
// -600
// -500
// -400
// -300
// -200
// -100
// 0
// 100
// 200
// 300
// 400
// 500
// 600
// 700
// 800
// 900
// 1000
// Ending inventory count =  1000
// Process 137818 has exited with status 0
// dlv dap (137772) exited with code: 0