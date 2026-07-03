```
***Question - The task is to read a log file, identify the types of log levels (INFO, ERROR, WARN), and count how many times each occurs.***


<!-- INFO server started
ERROR database connection failed
WARN disk usage high
INFO request processed
ERROR timeout -->

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main_() {
	countMap := make(map[string]int)

	file, err := os.Open("demo.log")

	defer file.Close()
	if err != nil {
		fmt.Println("file openning failed")
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) > 0 {
			word := parts[0]
			countMap[word]++
		}
	}

	for level, count := range countMap {
		fmt.Printf("%s\t%d\n", level, count)
	}

}

```