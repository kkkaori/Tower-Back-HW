package main

import (
	"fmt"
	"sync"
)

func sumSqrs(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sum := 0
	var wg sync.WaitGroup
	var lock sync.Mutex //mutex нужен для избежание гонки данных

	for _, val := range nums {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sqr := val * val
			lock.Lock() //сумму может изменять только одна горутина в одно время
			sum += sqr
			lock.Unlock()

		}(val)
	}

	wg.Wait()
	return sum
}

// все таки разобралась в каналах
func sumSqrs2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sum := 0
	ch := make(chan int, len(nums))

	go func() {
		for _, val := range nums {
			go func(val int) {
				sqr := val * val
				ch <- sqr
			}(val)
		}
	}()

	for i := 0; i < len(nums); i++ {
		sum += <-ch
	}

	close(ch)
	return sum
}

func main() {
	num := [5]int{2, 4, 6, 8, 10}
	sum := sumSqrs(num[:])
	//sum := sumSqrs2(num)

	fmt.Printf("Sum of squares = %d\n", sum)

}
