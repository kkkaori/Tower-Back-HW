package main

import (
	"fmt"
	"sync"
)

func process(input <-chan int, output chan<- int) {
	for val := range input {
		output <- val * 2
	}
	close(output)
}

func read(out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	var wg sync.WaitGroup

	input := make(chan int, 10)
	output := make(chan int, 10)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go process(input, output)
	wg.Add(1)
	go read(output, &wg)

	for _, val := range arr {
		input <- val
	}
	close(input)
	wg.Wait()

}
