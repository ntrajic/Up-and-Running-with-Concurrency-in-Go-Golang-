// Worker pools demo

package main

import (
	"fmt"
	"time"
)

func main() {

	const numJobs = 100                           	// We have this number of jobs to complete
	jobsChan := make(chan int, numJobs)          	// creates a buffered channel large enough to hold all jobs  NON_BLOCKING CHANNEL, w/ CAPACITY==#ofJobs (100)
	completedJobsChan := make(chan int, numJobs) 	// creates a buffered channel large enough to hold all completed jobs

	for w := 1; w <= 3; w++ { 						// this is number of workers.  Each worker will be a goroutine.  (3 go coros)
		go worker(w, jobsChan, completedJobsChan) 	// create worker w - pass worker number and both channels
	}

	for j := 1; j <= numJobs; j++ {
		jobsChan <- j 								// This loads the jobsChan channel with job numbers, PUSH ALL 100 JOBS to jobsChan
	}
	close(jobsChan) 								// Close jobsChan channel for input after all jobs have been loaded.  Channel must be closed in order to call "range" function.

	for a := 1; a <= numJobs; a++ {   				// READ completed jobs from completedJobsChan:
		<-completedJobsChan 						// Reads the completedJobsChan channel and does nothing with the contents.  
													// Point is to clear the channel and to delay termination of the program until all jobs are reported as finished.
	}
}

// worker shell for gocoro reads 100 jobs from jobChan and after simulated exectuion (sleep(2)) pushes completed job to the completedJobChan
func worker(id int, jobsChan <-chan int, completedJobsChan chan<- int) { // this syntax restricts the direction of each channel.  For THIS specific function, we will only SEND to completedJobsChan and RECEIVE from jobsChan

	for j := range jobsChan { 						// iterates (and RECEIVES) each and all the jobs in the channel.  Interesting that range seems to have its own receiver.
		fmt.Println("worker", id, "started  job", j, "with", len(jobsChan), "jobs left to process")
		time.Sleep(time.Second * 2) 				// simulates "work" that takes sleep time to complete
		fmt.Println("worker", id, "             finished job", j)
		completedJobsChan <- j 						// Loads finished job numbers into the completedJobsChan channel.
	}
}
// OUT: 
// Starting: /go/bin/dlv dap --listen=127.0.0.1:38237 --log-dest=3 from /workspaces/Up-and-Running-with-Concurrency-in-Go-Golang-/10-Concurrency_Patterns/5-workerPools
// DAP server listening at: 127.0.0.1:38237
// worker 3 started  job 1 with 99 jobs left to process
// worker 1 started  job 2 with 98 jobs left to process
// worker 2 started  job 3 with 97 jobs left to process
// worker 3              finished job 1
// worker 3 started  job 4 with 96 jobs left to process
// worker 1              finished job 2
// worker 2              finished job 3
// worker 2 started  job 5 with 95 jobs left to process
// worker 1 started  job 6 with 94 jobs left to process
// worker 2              finished job 5
// worker 2 started  job 7 with 93 jobs left to process
// worker 3              finished job 4
// worker 3 started  job 8 with 92 jobs left to process
// worker 1              finished job 6
// worker 1 started  job 9 with 91 jobs left to process
// ...
// worker 3              finished job 92
// worker 3 started  job 95 with 5 jobs left to process
// worker 2              finished job 91
// worker 2 started  job 96 with 4 jobs left to process
// worker 2              finished job 96
// worker 2 started  job 97 with 3 jobs left to process
// worker 3              finished job 95
// worker 3 started  job 98 with 2 jobs left to process
// worker 1              finished job 94
// worker 1 started  job 99 with 1 jobs left to process
// worker 1              finished job 99
// worker 1 started  job 100 with 0 jobs left to process
// worker 2              finished job 97
// worker 3              finished job 98
// worker 1              finished job 100
// Process 154458 has exited with status 0
// dlv dap (154414) exited with code: 0