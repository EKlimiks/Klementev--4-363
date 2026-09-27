package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup

	wait.Add(1)

	go func() {
		defer wait.Done()

		for number := 1; number <= 5; number++ {
			fmt.Println(number)
			time.Sleep(1 * time.Second)
		}
	}()

	wait.Wait()
	fmt.Println("Готово")
}
