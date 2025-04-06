package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

func print(str string) {
	var i int
	for i = 1; i <= 3; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(str, " ")
	}
}

var signal = []string{"test"}
var wg sync.WaitGroup // pointer
var mutex sync.Mutex  //pointer

func main_() {
	// =================Arrays=============================
	var arr [4]string
	arr[0] = "ram"
	arr[1] = "shyam"
	arr[3] = "pan"
	fmt.Println("array value is =", arr)
	fmt.Println("array Length is =", len(arr))
	// =================Slices=============================
	var name = []string{"ankit", "ashish", "manish", "bipin"}
	fmt.Println(name)
	name = append(name, "rajesh", "mukesh", "mahesh")
	fmt.Println("After Appending name is ", name)
	name = append(name[:5])
	fmt.Println("After slicing value is ", name)
	// make keyword allocate memory
	myScore := make([]int, 4)
	myScore[0] = 13
	myScore[1] = 14
	myScore[2] = 5
	myScore[3] = 63
	// MORE Memory appened from here
	myScore = append(myScore, 45, 56, 76, 5, 3, 2)
	fmt.Println("Using make Keyword add more values slices= ", myScore)
	sort.Ints(myScore)
	fmt.Println("Using build keyword sort values in slices= ", myScore)
	// how to remove a value from slices based on index valeus
	var course = []string{"react", "javascript", "python", "java", "go", "c++"}
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	// course =
	fmt.Println(course)
	// =================GOROUTINES=============================
	print("goroutine")
	print("SubGoroutine")
	var value *int
	// =================POINTER=============================
	var a int = 5
	b := &a
	fmt.Println("value of a is", a)
	fmt.Println("Pointer value is", b)
	fmt.Println("Pointer value is", *b)
	*b = *b + 8
	fmt.Println("value of a is", a)
	// =================Defer===============================
	defer fmt.Println("ashish")
	defer fmt.Println("jain")
	fmt.Println("world")
	fmt.Println("Hello")
	callDefer()
	// =================Wait Group===============================
	websiteList := []string{
		"https://google.com/",
		"https://leetcode.com/",
		"https://www.facebook.com/",
		"https://github.com/",
		"https://google.com/",
		"https://leetcode.com/",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signal)
	// =================Anonymous function and mutex(Lock and UnLock)===============================
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	mutR := &sync.RWMutex{}
	wg.Add(3)
	count := []int{}
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("First")
		mut.Lock()
		count = append(count, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Second")
		mut.Lock()
		count = append(count, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Third")
		mut.Lock()
		count = append(count, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	mutR.RLock()
	fmt.Println("Reading values of count =>>>", count)
	mutR.RUnlock()
	// =================Channel example===============================
	myChan := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// Receiving channel
	go func(wg *sync.WaitGroup, myChan <-chan int) { //  receiving value from channel
		// for i := 0; i < 5; i++ {
		val, isOpen := <-myChan
		if isOpen {
			fmt.Println("Received value is", val)
		} else {
			fmt.Println("Channel is closed")
		}
		// fmt.Println("Value getting from channel here ", <-myChan)
		// }
		wg.Done()
	}(wg, myChan)
	// Sending channel
	go func(wg *sync.WaitGroup, myChan chan<- int) { //  Sending value into channel
		// for i := 0; i < 5; i++ {
		myChan <- 8
		close(myChan)
		// myChan <- 1
		// myChan <- 9

		// }
		wg.Done()
	}(wg, myChan)
	// waiting from main function
	wg.Wait()
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("THere is error")
	} else {
		mutex.Lock()
		signal = append(signal, endpoint)
		mutex.Unlock()
		fmt.Printf("%d status code for website %s", res.StatusCode, endpoint)
		fmt.Println("")
	}
}

// ====================== Defer Function ================
func callDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("=>", i)
	}
}

// stack = ashish, jain, 0,1,2,3,4
// o/p = world, Hello, 4,3,2,1,0, jain, ashish
// ===========================================================
