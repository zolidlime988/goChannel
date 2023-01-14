package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"www.google.com",
		"www.facebook.com",
		"www.stackoverflow.com",
		"www.golang.org",
		"www.amazon.com",
	}

	c := make(chan string)
	for _, url := range urls {
		go checkLink(url, c)
	}
	for url := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkLink(url, c)
		}(url)
	}
}

func checkLink(url string, c chan string) {
	res, err := http.Get("http://" + url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}
	if res.StatusCode == 200 {
		fmt.Println(url, "is up")
		c <- url
	}

}
