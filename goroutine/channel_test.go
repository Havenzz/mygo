package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var (
	numChan = make(chan int, 2000)
	resChan = make(chan int, 2000)
	mu      sync.Mutex
	counter = 8
)

func writeNumChan() {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func writeResChan() {
	for v := range numChan {
		sum := 0
		for i := 1; i <= v; i++ {
			sum += i
		}
		resChan <- sum
	}
	mu.Lock()
	counter--
	if counter == 0 {
		close(resChan)
	}
	mu.Unlock()
}

func TestChannel(t *testing.T) {
	go writeNumChan()
	for i := 0; i < 8; i++ {
		go writeResChan()
	}

	i := 0
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Printf("res[%v] = %v\n", i, v)
		i++
	}

	for v := range make(chan int, 10) {
		fmt.Println("v = ", v)
	}
}
