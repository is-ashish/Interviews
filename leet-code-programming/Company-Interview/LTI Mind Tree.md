```
> concurrency vs parallelism
> How garbage collector done in Golang which Data structure Used here.
> Have you worked on CI/CD and Cloud platform
>
>

```
// write a programm handles 2 urls and response should be return into main go routines

// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	urls := []string{"https://www.google.com", "https://www.hotstar.com"}

	for _, url := range urls {
		wg.Add(1)
		go httpCall(url, &wg)
	}

	wg.Wait()
}

func httpCall(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("response", res)
	}
}


error -: 


error Get "https://www.hotstar.com": dial tcp: lookup www.hotstar.com on 169.254.169.254:53: dial udp 169.254.169.254:53: connect: no route to host
error Get "https://www.google.com": dial tcp: lookup www.google.com on 169.254.169.254:53: dial udp 169.254.169.254:53: connect: no route to host
