// This script test concurrent updates performance between using sync/atomic and sync.Mutext.
// run by command: go run main2.go
//
// The result shown below, they are almost the same.
//
//     Mutex  count: 1000000 Using: 561.377536ms
//     Atomic count: 1000000 Using: 539.912381ms
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count int64

const LOOP = 1000000

func useMutex() {
	worker := func(wait *sync.WaitGroup, mu *sync.Mutex) {
		mu.Lock()
		count += 1
		mu.Unlock()
		wait.Done()
	}

	mu := &sync.Mutex{}
	wait := &sync.WaitGroup{}
	for i := 0; i < LOOP; i++ {
		wait.Add(1)
		go worker(wait, mu)
	}
	wait.Wait()
}

func useAtomic() {
	worker := func(wait *sync.WaitGroup) {
		atomic.AddInt64(&count, 1)
		wait.Done()
	}

	wait := &sync.WaitGroup{}
	for i := 0; i < LOOP; i++ {
		wait.Add(1)
		go worker(wait)
	}
	wait.Wait()
}

func main() {
	count = 0
	start := time.Now()
	useMutex()
	fmt.Println("Mutex  count:", count, "Using:", time.Since(start))

	count = 0
	start = time.Now()
	useAtomic()
	fmt.Println("Atomic count:", count, "Using:", time.Since(start))
}
