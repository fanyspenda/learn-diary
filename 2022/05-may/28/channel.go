//unbuffered channel
// channel tanpa buffer - data yang dikirimkan pada channel harus diconsume
//terlebih dahulu sebelum dapat digunakan kembali.
package main

import (
	"fmt"
)

func main() {

	tambahChannel := make(chan int)

	// membuat goroutine/thread
	go tambah(tambahChannel)
	//mengirimkan data ke tambah() melalui channel
	tambahChannel <- 5
	//menerima dari channel setelah diproses oleh tambah()
	result := <-tambahChannel
	fmt.Println(result)

	//mengirimkan data ke tambah() melalui channel
	tambahChannel <- 1
	//menerima dari channel setelah diproses oleh tambah()
	result = <-tambahChannel
	fmt.Println(result)

	close(tambahChannel)
}

func tambah(channel chan int) {
	for i := range channel {
		channel <- i + i
	}
}
