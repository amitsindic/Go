package main

import (
	"fmt"
	"time"
)


func main() {
	go PrintNumbers([]int{1, 2, 3, 4, 5}, "Goroutine 1")
	go PrintNumbers([]int{6, 7, 8, 9, 10}, "Goroutine 2")
	time.Sleep(1 * time.Second)
}

func PrintNumbers(arr []int , goRoutineName string) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Printing from %s: %d\n", goRoutineName, arr[i])
	}
}