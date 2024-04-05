package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func badUsage() {
	wg := sync.WaitGroup{}
	var v uint64

	for i := 0; i < 3; i++ {
		go func() {
			wg.Add(1)
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(v)
}

func main() {
	wg := sync.WaitGroup{}
	var v uint64

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			v++
			wg.Done()
		}()
	}
}
