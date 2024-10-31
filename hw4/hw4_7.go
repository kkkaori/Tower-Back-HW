package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]int),
	}
}

func (sm *SafeMap) add(key int, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) get(key int) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.m[key]
	return value, ok

}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.add(i, i*i)
		}(i)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		value, ok := safeMap.get(i)
		if ok {
			fmt.Printf("Key: %d\n", value)
		}
	}

}
