package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazom.com",
	}

	//initial channels
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//get from channels
	for {
		i, open := <-c
		if !open {
			break
		}
		fmt.Println(i)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down"
		return
	}
	fmt.Println(link, "is up!")
	c <- "It is up"
}
