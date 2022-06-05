package main

import (
	"fmt"
	"sync"
)

const (
	totalWorker = 5
)

var lc = sync.Mutex{}

//goals: menambah 1 ke result setiap loop.
func main() {

	var (
		result     int
		wg         = sync.WaitGroup{}
		totalLoop  = 5000
		tambahChan = make(chan int, totalLoop)
	)

	wg.Add(totalWorker)
	for i := 0; i < totalWorker; i++ {
		go tambah(tambahChan, &result, &wg)
	}

	for i := 0; i < totalLoop; i++ {
		tambahChan <- 1
	}
	close(tambahChan)
	wg.Wait()
	fmt.Println(result)
}

func tambah(c chan int, result *int, wg *sync.WaitGroup) {
	for v := range c {
		lc.Lock() // --> menggunakan mutex karena terjadi race condition antar worker
		*result += v
		lc.Unlock()
	}
	wg.Done()
}
