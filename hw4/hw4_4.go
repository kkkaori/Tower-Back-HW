package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Workers struct {
	wg     sync.WaitGroup
	ch     chan int
	cancel context.CancelFunc
}

func newWorkers(amount int, cancel context.CancelFunc) *Workers {
	return &Workers{
		cancel: cancel,
		ch:     make(chan int, amount),
	}
}

func (w *Workers) start(N int, ctx context.Context) {
	for i := 0; i < N; i++ {
		w.wg.Add(1)
		go func(i int) {
			defer w.wg.Done()
			for {
				select {
				case data := <-w.ch:
					fmt.Printf("Workers %d data: Value %v\n", i, data)
				case <-ctx.Done():
					fmt.Println("Worker ", i, " stopped")
					return
				}
			}
		}(i)
	}
}

func (w *Workers) stop() {
	w.cancel()
	w.wg.Wait()
	close(w.ch)
}

func (w *Workers) send(data int) {
	w.ch <- data
}

func main() {
	//проверка ввода
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run hw4_4.go <number_of_workers>")
		return
	}
	N, err := strconv.Atoi(os.Args[1])
	if err != nil || N == 0 {
		fmt.Println("Please provide a valid number of workers(>1)")
		return
	}
	ctx, cncl := context.WithCancel(context.Background())

	workers := newWorkers(N, cncl)
	workers.start(N, ctx)

	//отслеживаем завершение программы
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//бесконечный цикл
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		num := gen.Intn(1000)
		select {
		case <-sigs:
			workers.stop()
			return
		default:
			workers.send(num)
		}
		time.Sleep(time.Millisecond * 100)
	}

}
