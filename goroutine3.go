package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.RWMutex) {
	m.Lock()
	x++
	m.Unlock()
	wg.Done()

}

func main() {
	var w = sync.WaitGroup{}
	var m = sync.RWMutex{}
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}

	w.Wait()
	fmt.Println("value of x", x)
}
