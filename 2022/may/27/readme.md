## Sinkronisasi Thread dengan Mutex

## Race condition
Kondisi dimana suatu proses saling bertabrakan dengan proses lainnya. Contoh:

Terdapat saldo sebesar 100.000 dan ada 2 Proses, proses menambah jumlah saldo dan mengurangi jumlah saldo. Ketika proses tersebut berjalan bersamaan, kedua proses tersebut saling mengganti nilai akhir dari saldo. Sehingga, mengakibatkan ketidak-konsistenan data.

misal, menambah 100.000 dan mengurangi 10.000. Ketika proses menambah selesai, proses tersebut akan memasukkan data saldo ke database sebesar 200.000. Namun, proses mengurangi juga akan mengganti data saldo tersebut menjadi 90.000.

Disinilah kita perlu menerapkan mutex :)

## Mutex (Mutual Exclusion Object)

mutex adalah sebuah mekanisme dimana kita dapat berbagi resource antar thread dengan kemampuan untuk membatasi resource tersebut agar tidak digunakan secara bersamaan.

Dari contoh di atas:

Ketika saldo ditambah, kita dapat menggunakan mutex untuk mengunci saldo, sehingga ketika proses pengurangan saldo masuk, ia akan menunggu (blocking) hingga proses penambahan saldo selesai. baru proses pengurangan saldo dapat berjalan. Begitu pula sebaliknya.


> **note:**
> blocking disini berbeda dengan mencegah agar program tidak finish. Block di sini artinya goroutine menunggu goroutine lain untuk menyelesaikan tugasnya terlebih dahulu. Jika kita tidak implementasikan waitGroup pada fungsi main misalnya, maka program akan finish tanpa sempat menyelesaikan goroutine yang masih menunggu.

**Pada Golang**, mutex dapat diterapkan dengan cara:

```go
mu sync.Mutex
func HitungSaldo (saldo int64) int64 {
    mu.Lock()
    //sebuah proses yang akan dikunci agar thread (goroutine) lain tidak bisa mengakses proses ini
    mu.Unlock()
}
```

## Readers-Writer Mutex

Pada implementasi di atas, mutex mencegah thread/goroutine lain mengakses fungsi `HitungSaldo` sepenuhnya. Dengan begitu, akan muncul blocking dimana jika `HitungSaldo` diakses oleh 10.000 goroutine, maka akan ada 9.999 goroutine yang menunggu 1 proses selesai.

Di sinilah readers-writer membantu meningkatkan performa goroutine dengan konsep **readers-writer mutex**.

readers-writer mutex bekerja dengan cara mengizinkan akses ke proses yang kita lock jika proses tersebut melakukan pembacaan data saja. Jika melakukan pengubahan data, maka proses tersebut akan terkunci (blocking). Dengan begitu, proses dapat berjalan lebih cepat.

Pada golang, implementasi Readers-Writer dapat menggunakan `sync.RWMutex`