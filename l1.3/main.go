package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Worker %d got %d\n", id, val)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run l1.3/main.go <num_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Введите правильное число воркеров")
		return
	}

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	go func() {
		counter := 0
		for {
			ch <- counter
			counter++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
	close(ch)
	wg.Wait()
}
