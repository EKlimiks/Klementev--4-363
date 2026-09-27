package main

import (
	"fmt"
	"time"
)

func search(name string, timeSleep time.Duration, result chan string) {
	time.Sleep(timeSleep)
	result <- name
}

func main() {
	result := make(chan string)

	go search("База 1", 3*time.Second, result)
	go search("База 2", 1*time.Second, result)
	go search("База 3", 2*time.Second, result)

	fmt.Println("Первый результат:", <-result)
}
