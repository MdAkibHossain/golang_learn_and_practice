package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signal []string
var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer
func main() {

	// go greeter("Hello")
	// greeter("World")
	var website = []string{
		"http://amazon.com",
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}
	for _, web := range website {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signal)
}

//	func greeter(s string) {
//		for i := 0; i < 6; i++ {
//			time.Sleep(3 * time.Millisecond)
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Opps in emdpoint")
	} else {
		mut.Lock()
		signal = append(signal, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
