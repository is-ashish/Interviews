```
> marshal and unmarshal object

> what is control plane

> what is history log

> what is kube-proxy and how it works

> roles and responsibility - daily activity in current project

> what is channels and wait groups

> what is goroutines and handle

> what gin framework do

> what is does migration and how you handled from GKE to EKS

> Produser- consumer problems

> Have you used concurrency in Golang

> What is worker Pool ? how it is works ? 


```


## Coding round


# 1.

```

Input should be -: 

{
    "name":"John Doe",
    "email”:”jd@email.io”,
    "password":"xxxxx",
    "phone_number": "+91907xxxxxx",
    "test": {
        "id":"1234",
        "tel": {
            "tags":["hello", "world"]
        }
    }
}

Output: Flatten formate -:
 {
  "name": “John Doe",
  "email": “jd@email.io”,
  "password": "xxxxx",
  "phone_number": "+91907xxxxxx",
  "test_id”:”1234",
  "test_tel_tags_0": "hello",
  "test_tel_tags_1": "world"
 }
```

*** sample program for flatten object -: ***
```go
package main

import "fmt"

func main() {
    var a interface{}
    a = 42 // a holds an integer
    fmt.Printf("Type: %T, Value: %v\n", a, a)
    a = "hello" // a holds a string
    fmt.Printf("Type: %T, Value: %v\n", a, a)

}

```


*** print odd numbers from 0-n, using multiple worker go routines ***

# 2. 

```go
package main

import "fmt"

func main() {
    n := 100

    for i := 0; i <= n; i++ {
        go checkOdd(i)
    }

}

func checkOdd(i int) {
    if i%2 == 1 {
        fmt.Println(i)
    }
}
```


