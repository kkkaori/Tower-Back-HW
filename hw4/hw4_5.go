package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func read(ch chan int) {
	k := 1
	for data := range ch {
		fmt.Println("read data: ", data, " ", k, " time(s)")
		k++
	}
	fmt.Println("channel closed")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run hw4_5.go <time of work>")
		return
	}
	t, err := strconv.Atoi(os.Args[1])
	if err != nil || t == 0 {
		fmt.Println("Please provide a valid number of seconds(>1)")
		return
	}
	ch := make(chan int)
	go read(ch)

	dur := time.Second * time.Duration(t)
	timer := time.NewTimer(dur)

	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		num := gen.Intn(1000)
		select {
		case <-timer.C:
			close(ch)
			fmt.Println("time's out")
			return
		default:
			ch <- num
		}
		time.Sleep(time.Millisecond * 100)
	}

}
