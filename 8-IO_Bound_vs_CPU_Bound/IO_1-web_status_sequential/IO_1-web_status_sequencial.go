package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)



func main() {
	runtime.GOMAXPROCS(16)  // Extra processors ****don't help with sequential tasks****
	fmt.Println(runtime.NumCPU())

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


	start := time.Now()

	for _, link := range links {
		checkLink(link)
	}






	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLink(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")

		return
	}

	fmt.Println(link, "is LIVE!")

}
//OUT:
//Starting: /go/bin/dlv dap --listen=127.0.0.1:41969 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/8-IO_Bound_vs_CPU_Bound/IO_1-web_status_sequential
// DAP server listening at: 127.0.0.1:41969
// 2
// http://hashnode.com is LIVE!
// http://dev.to is LIVE!
// http://stackoverflow.com is LIVE!
// http://golang.org is LIVE!
// http://medium.com is LIVE!
// http://github.com is LIVE!
// http://techcrunch.com is LIVE!
// http://techrepublic.com is LIVE!
// Processes took 1.431780558s                             <---- 16 processors sequentially check 8 web_links for 1.4 sec 1st time, when cached below 1 sec.
// Process 108389 has exited with status 0
// dlv dap (108331) exited with code: 0