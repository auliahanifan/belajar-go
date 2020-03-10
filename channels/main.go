package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// make channel and initiate the type, we use string
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative Loop that easy to understand
	for l := range c {
		// literal function
		go func(link string) {
			time.Sleep(1 * time.Second)
			checkLink(link, c)
		}(l) // add l as arguments so, the function can run with the truly value of l
		// go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down!"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
	// c <- "Yes its up!"
}
