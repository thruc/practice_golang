package main

import "sync"

func parallelMergesortV1(a []int) {
	if len(a) <= 1 {
		return
	}

	mid := len(a) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		parallelMergesortV1(a[:mid])
		wg.Done()
	}()

	go func() {
		parallelMergesortV1(a[mid:])
		wg.Done()
	}()

	wg.Wait()

	merge(a, mid)
}

func merge(a []int, mid int) {
	tmp := make([]int, len(a))
	copy(tmp, a)

	i, j := 0, mid
	for k := 0; k < len(a); k++ {
		if i >= mid {
			a[k] = tmp[j]
			j++
		} else if j >= len(a) {
			a[k] = tmp[i]
			i++
		} else if tmp[j] < tmp[i] {
			a[k] = tmp[j]
			j++
		} else {
			a[k] = tmp[i]
			i++
		}
	}
}
