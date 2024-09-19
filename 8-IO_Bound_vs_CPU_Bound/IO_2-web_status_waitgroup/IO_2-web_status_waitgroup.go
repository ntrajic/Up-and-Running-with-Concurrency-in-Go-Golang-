package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(16) // Even ONE processor can take advantage of concurrency with IO bound code. More may not be needed, but do help.

	links := []string{
		"http://hashnode.com",
		"http://dev.to",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://medium.com",
		"http://github.com",
		"http://techcrunch.com",
		"http://techrepublic.com",
	}
	wg.Add(len(links))


	start := time.Now()

	for _, link := range links {
		go checkLink(link)
	}
	wg.Wait()






	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLink(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		wg.Done()
		return
	}

	fmt.Println(link, "is LIVE!")
	wg.Done()
}
// OUT:
// Starting: /go/bin/dlv dap --listen=127.0.0.1:46145 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/8-IO_Bound_vs_CPU_Bound/IO_2-web_status_waitgroup
// DAP server listening at: 127.0.0.1:46145
// http://techcrunch.com is LIVE!
// http://medium.com is LIVE!
// http://github.com is LIVE!
// http://techrepublic.com is LIVE!
// http://dev.to is LIVE!
// http://hashnode.com is LIVE!
// http://stackoverflow.com is LIVE!
// http://golang.org is LIVE!
// Processes took 496.901597ms    // for N=1 processor, for runtime.GOMAXPROCS(16) -> multiple processors (16) DO HELP w/ coros parallel exe: Processes took 396.40014ms
// Process 111245 has exited with status 0
// dlv dap (111186) exited with code: 0