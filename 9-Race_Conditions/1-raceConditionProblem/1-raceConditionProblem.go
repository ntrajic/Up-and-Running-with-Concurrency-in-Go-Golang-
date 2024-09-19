package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup

	widgetInventory int32= 1000  //Package-level variable will avoid all the pointers, NON-PROTECTED SHARED VAR, BUGGY!!!!
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
	for i := 0; i < 3000; i++ {

		widgetInventory -= 100

	}
	wg.Done()
}


func newPurchases() {  // 1000000 widgets purchased
	for i := 0; i < 3000; i++ {

		widgetInventory+= 100

	}
	wg.Done()
}
//OUT: UNREALIABLE BUGGY CODE!!
//Starting: /go/bin/dlv dap --listen=127.0.0.1:37643 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/9-Race_Conditions/1-raceConditionProblem
// DAP server listening at: 127.0.0.1:37643
// Starting inventory count =  1000
// Ending inventory count =  -2514                  <----- most time is 1000 (300 sold, 300 bought), but is unreliable, unprotected shared var
// Process 133795 has exited with status 0
// dlv dap (133754) exited with code: 0