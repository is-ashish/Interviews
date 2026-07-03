package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"
)

func main() {
	var mx sync.Mutex
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	wordCount := make(map[string]int)
	files := []string{"location.txt", "user.txt", "vehicle.txt"}
	go processFiles(ch1, files[0])
	go processFiles(ch2, files[0])
	go processFiles(ch3, files[2])

	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				mx.Lock()
				wordCount[v]++
				mx.Unlock()
			}
		case v1, ok := <-ch2:

			if !ok {
				ch2 = nil
			} else {
				mx.Lock()
				wordCount[v1]++
				mx.Unlock()
			}
		case v2, ok := <-ch3:
			if !ok {
				ch3 = nil
			} else {
				mx.Lock()
				wordCount[v2]++
				mx.Unlock()
			}
		}

		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}
	}

	type wordObject struct {
		word  string
		count int
	}
	var freqs []wordObject

	for k, v := range wordCount {
		freqs = append(freqs, wordObject{word: k, count: v})
	}

	// Sort by frequency (descending)
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	// fmt.Println(freqs)
	for _, v := range freqs {
		fmt.Printf("word -: %s ------- count -: %d \n", v.word, v.count)
	}
}

func processFiles(cha chan<- string, txt string) {

	file, err := os.Open(txt)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Change split function to scan words
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		cha <- word
	}
	close(cha)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
