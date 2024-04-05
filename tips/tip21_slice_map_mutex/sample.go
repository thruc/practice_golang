package main

import "sync"

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balances[id] += balance
}

func (c *Cache) AverageBalance() float64 {
	c.mu.RLock()
	balances := c.balances
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}
	return sum / float64(len(balances))
}
