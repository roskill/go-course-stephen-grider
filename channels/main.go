package main

import (
	"fmt"
	"net/http"
	"time"
)

var links = []string{
	"http://www.amazon.com",
	"http://www.golang.org",
	"http://www.stackoverflow.com",
	"http://www.mozilla.org",
	"http://www.google.com",
}

var channel = make(chan string)

func main() {
	// link are visited synchronously as Get is a blocking call
	for _, link := range links {
		go checkLink(link, channel)
	}

	// for i := 0; i < len(links); i++ {
	// fmt.Println(<-channel)
	// the first time the for loops runs, it blocks at Println,
	// then continues when the channel receives the first message
	// }

	// for { // infinite loop
	// 	go checkLink(<-channel, channel)
	// }

	for link := range channel { // slightly more intuitive
		// go checkLink(link, channel)
		go func(l string) { // pass value of link as l not reference
			time.Sleep(time.Second * 3)
			checkLink(l, channel)
		}(link)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		channel <- link
		// channel <- "Channel: Site is down"
		return
	}

	fmt.Println(link, "is up")
	channel <- link
	// channel <- "Channel: Site is up"
}
