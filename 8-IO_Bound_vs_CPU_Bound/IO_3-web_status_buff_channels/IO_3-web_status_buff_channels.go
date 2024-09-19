package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)



func main() {
	runtime.GOMAXPROCS(16)

	links := []string{
		"http://hashnode.com",
		"http://dev.to",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://medium.com",
		"http://github.com",
		"http://techcrunch.com",
	}
	c := make(chan string, len(links)) // NONBLOCKING create ***buffered channel**** large enough to hold all links


	start := time.Now()

	for _, link := range links {			// len(links)==7, range from 0 to 6
		go checkLink(link, c)
	}
	for len(c) < len(links) { 				// Infinite Loop ends when everything is in BUFFERED channel: loop while all links are checked in pparallel!!!

	}
	//for range links{  					// This also works beautifully, using blocking code
	//	fmt.Println("channel message:",<-c)
	//}

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		c <- link
		return
	}

	fmt.Println(link, "is LIVE!")
	c <- link
}
//OUT:
// //Starting: /go/bin/dlv dap --listen=127.0.0.1:36087 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/8-IO_Bound_vs_CPU_Bound/IO_3-web_status_buff_channels
// DAP server listening at: 127.0.0.1:36087
// http://dev.to is LIVE!
// http://techcrunch.com is LIVE!
// http://medium.com is LIVE!
// http://hashnode.com is LIVE!
// http://github.com is LIVE!
// http://stackoverflow.com is LIVE!
// http://golang.org is LIVE!
// Processes took 282.347966ms                        <---- super fast
// Process 114619 has exited with status 0
// dlv dap (114568) exited with code: 0