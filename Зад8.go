package main

import "fmt"

func main() {
	servers := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	for i := 0; i < 3; i++ {
		go func(n int) {
			for {
				request := <-servers[n]
				fmt.Println("Сервер", n+1, "обработал запрос", request)
			}
		}(i)
	}

	for i := 1; i <= 10; i++ {
		server := (i - 1) % 3
		servers[server] <- i
	}
}
