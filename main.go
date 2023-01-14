package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"www.google.com",
		"www.facebook.com",
		"www.stackoverflow.com",
		"www.golang.org",
		"www.amazon.comsss",
	}
	c := make(chan string)
	for _, url := range urls {
		go checkLink(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(url string, c chan string) {
	res, err := http.Get("http://" + url)
	if err != nil {
		c <- url + " Might Be Down"
		return
	}
	if res.StatusCode == 200 {
		//status[url] = true
		c <- url + " is! Up"
	}

}
