package main

import (
	"fmt"
	"sync"
)

func Squares(nums []int) []int {

	if len(nums) < 1 {
		return []int{}
	}

	sqrs := make([]int, len(nums))
	var wg sync.WaitGroup

	for i, val := range nums {
		wg.Add(1)
		go func(i, val int) {
			defer wg.Done()
			sqrs[i] = val * val
		}(i, val)
	}

	wg.Wait()
	return sqrs
}

func main() {
	num := [5]int{2, 4, 6, 8, 10}

	sqrs := Squares(num[:])

	for i := range sqrs {
		fmt.Printf("Square of %d is %d\n", num[i], sqrs[i])
	}

}
