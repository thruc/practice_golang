package main

import "sync"

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{
		counters: make(map[string]int),
	}
}

func (c Counter) BadIncrement(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func (c *Counter) GoodIncrement(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	counter := NewCounter()

	go func() {
		for i := 0; i < 10; i++ {
			counter.BadIncrement("foo")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			counter.BadIncrement("bar")
		}
	}()
}
