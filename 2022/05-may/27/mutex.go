package main

import (
	"fmt"
	"sync"
	"time"
)

var mu = sync.Mutex{}

func main() {

	var value = 100

	//secara logika, jika kita menambah 1000 dan mengurangi 1000,
	// maka hasilnya akan sama.
	//namun, tanpa menggunakan mutex, hasilnya akan berbeda-beda setiap running.
	for i := 0; i < 1000; i++ {
		go decrease(&value)
	}

	for i := 0; i < 1000; i++ {
		go increase(&value)
	}

	time.Sleep(3 * time.Second)

	fmt.Println(value)
}

//uncomment mutex untuk melihat contoh kasus anehnya.
func decrease(a *int) {
	// mu.Lock()
	*a = *a - 1
	// mu.Unlock()
}

func increase(a *int) {
	// mu.Lock()
	*a = *a + 1
	// mu.Unlock()
}
