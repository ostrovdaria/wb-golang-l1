package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("Воркер %d работает...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	fmt.Println("Нажмите Ctrl+c для завершения...")
	<-ctx.Done()
	fmt.Println("Программа завершилась.")
}
