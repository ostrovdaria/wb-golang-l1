package main

import (
	"fmt"
	"time"
)

func sender(ch chan int, stop chan bool) {
	i := 1
	for {
		select {
		case <-stop:
			close(ch)
			return
		default:
			ch <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}

	}
}

func receiver(ch chan int, done chan bool) {
	for val := range ch {
		fmt.Println(val)
	}
	done <- true
}

func main() {
	ch := make(chan int)
	stop := make(chan bool)
	done := make(chan bool)

	go sender(ch, stop)
	go receiver(ch, done)

	N := 5
	timer := time.After(time.Duration(N) * time.Second)

	<-timer
	stop <- true
	<-done
}
