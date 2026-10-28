<!-- QUESTION - et’s extend your program into a Gin HTTP server that exposes a GET endpoint to fetch the topper (highest scoring student across all subjects). -->
<!-- 
This is physics score card 

101,Vijay,physiscs,56,100
102,Raj,physiscs,98,100
103,Anil,physiscs,54,100

This is math 
102,Raj,Maths,90,100
101,Vijay,Maths,75,100
103,Anil,Maths,65,100

this is chemestry _

102,Raj,Chemistry,81,100
101,Vijay,Chemistry,76,100
103,Anil,Chemistry,84,100 -->

```
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var mu sync.Mutex

type score struct {
	ID   int `json:"id"`
	Mark int `json:"mark"`
}

func main_() {
	r := gin.Default()

	// Define get enpoint
	r.GET("/topper", func(c *gin.Context) {
		topStudent := getTopper()

		c.JSON(200, gin.H{
			"student_id":  topStudent.ID,
			"total_marks": topStudent.Mark,
			"message":     fmt.Sprintf("Topper is student %d with %d marks", topStudent.ID, topStudent.Mark),
		})
	})

	// Run server on localhost:8080
	r.Run(":8080")
}

func getTopper() score {

	wg := sync.WaitGroup{}

	wordCount := make(map[int]int)
	subjectFiles := []string{"physics.csv", "math.csv", "chemistry.csv"}

	for _, file := range subjectFiles {
		wg.Add(1)
		go processFile(file, &wg, wordCount)
	}
	wg.Wait()

	var result []score

	for k, v := range wordCount {
		result = append(result, score{ID: k, Mark: v})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Mark > result[j].Mark
	})

	return result[0]

	// fmt.Printf("Highest marks of student ID - %d and Total Marks %d\n", result[0].id, result[0].mark)

}

func processFile(file string, wg *sync.WaitGroup, wordCount map[int]int) {

	defer wg.Done()

	f, err := os.Open(file)

	if err != nil {
		fmt.Println("file process failed")
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()
		studentRecord := strings.Split(line, ",")

		studID, _ := strconv.Atoi(studentRecord[0])
		studMark, _ := strconv.Atoi(studentRecord[3])

		mu.Lock()
		wordCount[studID] = wordCount[studID] + studMark
		mu.Unlock()

	}
}


```