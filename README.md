


# Up-and-Running-with-Concurrency-in-Go-Golang-
Up and Running with Concurrency in Go (Golang), by Packt Publishing

IO bound vs CPU bound processes:
---------------------------------
* multiple cores do not help with sequential tasks
* just one core can see significant speed improvements with concurrency in IO-bound code (API calls) due to latency
* for CPU bond code, concurrency is no help with one core. CPU bound code sees significant improvements with inccrease of # of cores, up to a certain number of cores.
