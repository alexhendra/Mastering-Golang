package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func sumAtomic(from, to int, wg *sync.WaitGroup, res *int32, itemName string) {
	log.Printf(itemName)
	for i := from; i <= to; i++ {
		atomic.AddInt32(res, int32(i))
	}

	wg.Done()
	return
}

func sum(from, to int, wg *sync.WaitGroup, res *int, itemName string) {
	log.Printf(itemName)
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}
	wg.Done()
	return
}

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}
	wg.Add(4)

	// go sum(1, 25, wg, &s1, "Item 1")
	// go sum(26, 50, wg, &s1, "Item 2")
	// go sum(51, 75, wg, &s1, "Item 3")
	// go sum(76, 100, wg, &s1, "Item 4")

	go sumAtomic(1, 25, wg, &s1, "Item 1")
	go sumAtomic(26, 50, wg, &s1, "Item 2")
	go sumAtomic(51, 75, wg, &s1, "Item 3")
	go sumAtomic(76, 100, wg, &s1, "Item 4")

	wg.Wait()
	log.Println(s1)
}
