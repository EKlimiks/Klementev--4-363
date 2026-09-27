package main

import (
	"fmt"
	"net/http"
	"sync"
)

func check(tasks <-chan string, wait *sync.WaitGroup) {
	defer wait.Done()

	for address := range tasks {
		response, err := http.Get(address)

		if err != nil {
			fmt.Println("Ошибка:", address)
			continue
		}

		fmt.Println(address, response.Status)

		response.Body.Close()
	}
}

func main() {
	addresses := []string{
		"https://2gis.ru/?m=64.64699%2C54.566729%2F8.19",
		"https://music.yandex.ru/?from=iframe",
		"https://vk.ru/feed",
		"https://www.wikipedia.org",
		"https://workspace.ktk-45.ru/",
	}

	tasks := make(chan string)

	var wait sync.WaitGroup

	for number := 0; number < 3; number++ {
		wait.Add(1)
		go check(tasks, &wait)
	}

	for _, address := range addresses {
		tasks <- address
	}

	close(tasks)

	wait.Wait()

	fmt.Println("Проверка закончена")
}
