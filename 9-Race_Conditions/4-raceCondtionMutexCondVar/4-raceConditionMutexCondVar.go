package main

import (
	"fmt"
	"sync"
	)

var (
	wg sync.WaitGroup							//wg := 0
	mutex = sync.Mutex{}
	widgetInventory int32= 1000  				//Package-level variable will avoid all the pointers SHARED VAR
	newPurchase = sync.NewCond(&mutex)			//CONDITION
)


func main() {
	fmt.Println("Starting inventory count = ", widgetInventory)
	wg.Add(2)	// wg = 2
	go makeSales()
	go newPurchases()
	wg.Wait()	// wait for wg  == 0, wait for both coros to complete
	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales() {  // 1000000 widgets sold
	for i := 0; i < 3000; i++ {
		mutex.Lock()							//<<<<<
		if widgetInventory-100 < 0{
			newPurchase.Wait()					//WAIT when there's nothing on the stack to sell  <------  on signal wake up from waiting
		}										//                                                       |
		widgetInventory -= 100                  //                                                       1
		fmt.Println(widgetInventory)            //                                                       |
		mutex.Unlock()							//<<<<<                                              signal that single purchase is doen
	}

	wg.Done() // ALL 3000 SALES ARE DONE (wg -= 1)
}

func newPurchases() {  // 1000000 widgets purchased
	for i := 0; i < 3000; i++ {
		mutex.Lock()							//<<<<<                1
		widgetInventory+= 100					//                     ^
		fmt.Println(widgetInventory)			//                     |
		newPurchase.Signal()					// SIGNAL --------------
		mutex.Unlock()							//<<<<<
	}
	wg.Done()	// ALL 3000 PURCHASES ARE DONE (WG -= 1)
}
// OUT:
// 40200
// 40300
// 40400
// 40500
// 40600
// 40700
// 40800
// 40900
// 41000
// 41100
// 41200
// 41300
// 41400
// 41500
// 41600
// 41700
// 41800
// 41900
// 42000
// 42100
// 42200
// 42300
// 42400
// 42500
// 42600
// 42700
// 42800
// 42900
// 43000
// 43100
// 43200
// 43300
// 43400
// 43500
// 43600
// 43700
// 43800
// 43900
// 44000
// 44100
// 44200
// 44300
// 44400
// 44500
// 44600
// 44700
// 44800
// 44900
// 45000
// 45100
// 45200
// 45300
// 45400
// 45500
// 45600
// 45700
// 45800
// 45900
// 46000
// 46100
// 46200
// 46300
// 46400
// 46500
// 46600
// 46700
// 46800
// 46900
// 47000
// 47100
// 47200
// 47300
// 47400
// 47500
// 47600
// 47700
// 47800
// 47900
// 48000
// 48100
// 48200
// 48300
// 48400
// 48500
// 48600
// 48700
// 48800
// 48900
// 49000
// 49100
// 49200
// 49300
// 49400
// 49500
// 49600
// 49700
// 49800
// 49900
// 50000
// 50100
// 50200
// 50300
// 50400
// 50500
// 50600
// 50700
// 50800
// 50900
// 51000
// 51100
// 51200
// 51300
// 51400
// 51500
// 51600
// 51700
// 51800
// 51900
// 52000
// 52100
// 52200
// 52300
// 52400
// 52500
// 52600
// 52700
// 52800
// 52900
// 53000
// 53100
// 53200
// 53300
// 53400
// 53500
// 53600
// 53700
// 53800
// ...
// 800
// 1700
// 1600
// 1500
// 1400
// 1300
// 1200
// 1100
// 1000
// Ending inventory count =  1000                           <<<<< 3000 sold, 3000 purchased, starting # of products, inventory = 1000 => no change
// Process 147514 has exited with status 0