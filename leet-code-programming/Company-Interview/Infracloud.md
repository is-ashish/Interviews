

***Assignment***
1. Build a URL Sortner with analytics and deploy over docker
2. Sample program to dry run give output on console -: based on pointer

``` go
> Assignment based prgramming
    > What if all charSET used in shortIT 
    > If short ID is already exits
    > At the same time 2 user is trying to sorten URL (If User is trying to shorten same URL) ? which error come ?
> What is Kube proxy in k8s ?
> what are the features covered through Unit test
> Production bug fixes ? In recent which issues fixed

```

*** program give output- Answer= 6 *** 
```go
package main

import (
	"fmt"
)

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c Counter) Get() int {
	return c.count
}

func (c Counter) Decrement() {
	c.count--
}

func main() {
	c := &Counter{4}
	c.Increment()
	c.Decrement()
	c.Decrement()
	c.Increment()
	c.Decrement()
	fmt.Println(c.Get())
}
 ====
```