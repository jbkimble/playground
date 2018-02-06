package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}


	// using a range with a channel
	// wait for the channel to return some value, after it has returned a value 
	// assign it to the variable 'l' then run the body of the for loop
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

	// for {
	// 	// value coming through 'c' is a blocking call
	// 	go checkLink(<-c, c)
	// }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}


// channel <- 5 // send the value '5' into this channel
// myNumber <- channel // wait for a value to be sent into the channel.  When we get one, assign the value to 'myNumber'
// fmt.Println(<- channel) // Wait for a value to be sent into the channel.  When we get one, log it out immediately