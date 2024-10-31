package main

import (
	"fmt"
)

func hasBit(num int64, pos uint) bool {
	val := num & (1 << pos)
	return val > 0
}

func ChangeNBit(num int64, pos uint) int64 {
	pos--
	if hasBit(num, pos) {
		num &= ^(1 << pos)
	} else {
		num |= (1 << pos)
	}
	return num
}

func main() {
	num := int64(64)
	fmt.Printf("Number: %d\n in bytes: %08b\n", num, num)

	changedNum := ChangeNBit(num, 1)
	fmt.Printf("Number: %d\n in bytes:  %08b\n", changedNum, changedNum)
}
