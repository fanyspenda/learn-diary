package main

import (
	"fmt"
	"sync"
	"time"
)

var mu = sync.Mutex{}
var kurangCond = sync.NewCond(&mu)

func main() {

	var value = 0

	for i := 0; i < 1000; i++ {
		go decrease(&value)
	}

	for i := 0; i < 1000; i++ {
		go increase(&value)
	}

	time.Sleep(2 * time.Second)

	fmt.Println(value)
}

//uncomment mutex untuk melihat contoh kasus anehnya.
func decrease(a *int) {
	mu.Lock()
	for *a-2 < 0 {
		kurangCond.Wait()
	}
	*a = *a - 2
	mu.Unlock()
}

func increase(a *int) {
	mu.Lock()
	*a = *a + 1
	kurangCond.Signal()
	mu.Unlock()
}
