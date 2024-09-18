package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	start := time.Now()
	for i := 1; i < 50; i++ {

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
		fmt.Println("\n", responseObject.Joke)

	}
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}
//OUT: 
//
//  Remember, the best angle to approach a problem from is the "try" angle.
//
//  What do you call an old snowman? Water.
//
//  Doctor: Do you want to hear the good news or the bad news?
// Patient: Good news please.
// Doctor: we're naming a disease after you.
//
//  I invented a new word! Plagiarism!
//
//  Why did the clown have neck pain? - Because he slept funny
// Processes took 5.316473873s                                              <-------
// Process 12695 has exited with status 0
// dlv dap (11343) exited with code: 0

