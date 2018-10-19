// START1 OMIT

package main

import (
	"fmt"
	"net/http"
)

func getStatusCode(url string, ch chan<- int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	ch <- resp.StatusCode
}

// STOP1 OMIT

// START2 OMIT

func main() {
	ch := make(chan int)

	urls := []string{
		"https://facebook.com",
		"https://twitter.com",
		"https://github.com/",
	}

	for _, u := range urls {
		go getStatusCode(u, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}

// STOP2 OMIT