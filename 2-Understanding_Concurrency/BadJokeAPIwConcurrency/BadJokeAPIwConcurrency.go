package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	wg.Add(50)
	for i := 1; i <= 50; i++ {
		go getBadJoke()
}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}
func getBadJoke(){
		client := &http.Client{}
		request, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
		if err != nil {
			fmt.Print(err.Error())
		}
		request.Header.Add("Accept", "application/json")
		request.Header.Add("Content-Type", "application/json")

		response, err := client.Do(request)

		if err != nil {
			fmt.Print(err.Error())
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(response.Body)

		bodyBytes, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Print(err.Error())
		}
		var responseObject Response
		err = json.Unmarshal(bodyBytes, &responseObject)
		if err != nil {
			return 
		}
		//fmt.Println("\nresponse", response)
		//fmt.Println("\nresponse.Body", response.Body)
		//fmt.Println("\nbodyBytes", bodyBytes)
		//fmt.Println("\nresponseObject", responseObject)

		fmt.Println("\n",responseObject.Joke)
		wg.Done()
	}


//OUT:
//
//  What did the piece of bread say to the knife? Butter me up.
//
//  The rotation of earth really makes my day.
//
//  What did the left eye say to the right eye? Between us, something smells!
//
//  Sometimes I tuck my knees into my chest and lean forward.  That’s just how I roll.
//
//  Every night at 11:11, I make a wish that someone will come fix my broken clock.
//
//  I tried to milk a cow today, but was unsuccessful. Udder failure.
// Processes took 1.085142267s                                           <-------
// Process 17787 has exited with status 0
// dlv dap (17736) exited with code: 0

