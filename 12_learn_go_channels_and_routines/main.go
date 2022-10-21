package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLinks(link, c)
	}
	// for i := 0; i < len(links); i++ {     /// s
	// 	fmt.Println(<-c)
	// }
	/////
	// for {
	// 	go checkLinks(<-c, c)
	// }
	// ----------OR--------
	// for l := range c {
	// 	go checkLinks(l, c)
	// }
	/////
	for l := range c {
		go func(link string) {
			time.Sleep(1 * time.Second)
			checkLinks(link, c)
		}(l)
	}
}
func checkLinks(link string, c chan<- string) { //  chan<- means :: send only
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Down!"    ///  s
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	//c <- "Up!"   ///   s
	c <- link
}
