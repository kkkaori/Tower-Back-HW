package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func work() int {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := gen.Intn(1000)
	return num
}

func end1(workCh chan int) {
	end := make(chan struct{}) //struct чтобы не занимать память

	go func() {
		for {
			select {
			case workCh <- work():
				fmt.Println("Sending data ...")
			case <-end:
				close(workCh)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		time.Sleep(time.Second * 1)
		end <- struct{}{}
	}()

	for data := range workCh {
		fmt.Printf("Read data: %d\n", data)
	}

	fmt.Printf("1 end")
}

func end2(ch chan int) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case ch <- work():
				fmt.Println("Sending data ...")
			case <-ctx.Done():
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}(ctx)

	go func() {
		time.Sleep(time.Second * 1)
		cancel()
	}()

	for data := range ch {
		fmt.Printf("Read data:  %d\n", data)
	}

	fmt.Printf("2 end")

}

func main() {

	workCh := make(chan int)

	// by channel
	//end1(workCh)

	// by context
	end2(workCh)

}
