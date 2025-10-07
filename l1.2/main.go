package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num * num)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, n := range numbers {
		wg.Add(1)
		go square(n, &wg)
	}

	wg.Wait()
}
