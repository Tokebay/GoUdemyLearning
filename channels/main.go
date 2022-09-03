package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{

		"https://halykbank.kz",
		"https://homebank.kz/",
		"https://google.kz/",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// fmt.Println(<-c) // main routine

	for link := range c {
		// fmt.Println(i, <-c)
		time.Sleep((5 * time.Second))
		go checkLink(link, c)
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, ": is down")
		c <- "might be down!"
		return
	}

	fmt.Println(link, "is up!")
	c <- "it is up" // go routine

}
