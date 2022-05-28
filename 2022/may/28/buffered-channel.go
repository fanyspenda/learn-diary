//unbuffered channel
// channel tanpa buffer - data yang dikirimkan pada channel harus diconsume
//terlebih dahulu sebelum dapat digunakan kembali.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var total int

func main() {

	tambahChannel := make(chan int, 9)
	go tambah(tambahChannel)

	wg.Add(20)
	for i := 0; i < 20; i++ {
		fmt.Println("kirim data")
		tambahChannel <- 1
	}

	//waitGroup diperlukan untuk mencegah proses dibawah berjalan
	//ketika goroutine masih memproses perhitungan fungsi tambah()
	wg.Wait()
	result := total
	fmt.Println(result)

	close(tambahChannel)
}

func tambah(channel chan int) {
	for n := range channel {
		total += n
		fmt.Println("terima data")
		wg.Done()
	}
}

//bila dilihat hasilnya, data yang dikirimkan adalah 10 data.
// setelah 10 data terkirim, goroutine baru menjalankan fungsi tambah().
