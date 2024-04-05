package main

import "sync/atomic"

func main() {
	var i int64

	go func() {
		atomic.AddInt64(&i, 1)
	}()

	go func() {
		atomic.AddInt64(&i, 1)
	}()
}

func channel() {
	var i int64
	c := make(chan int64)

	go func() {
		c <- 1
	}()

	go func() {
		c <- 1
	}()

	i += <-c
	i += <-c
}
