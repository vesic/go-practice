package main

// Example program that uses goroutines and channels.
// It fetches several web pages simultaneously using the net/http package,
// and prints the URL with most bytes in the response.

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}

	results := make(chan Page)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- Page{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var page Page

	for range urls {
		result := <-results
		if result.Size > page.Size {
			page = result
		}
	}

	fmt.Println("Most bytes page:", page.URL)
}
