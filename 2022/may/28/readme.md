## Message Passing

Message passing adalah salahsatu dari 2 cara untuk berkomunikasi antar thread.

Pada materi sebelumnya, dijelaskan bahwa berkomunikasi antar thread dibagi menjadi `shared memory` dan `message passing`. 

`shared memory` memerlukan handling agar tidak ada proses yang mengubah data pada 1 memori yang sama.

Sedangkan pada `message passing`, kita berkomunikasi antar thread menggunakan channel (pada Golang).

## Channel pada Golang

Channel merupakan sebuah medium antar goroutine/thread untuk mengirimkan data pada/dari goroutine. Channel bersifat blocking jika:
1. Channel mengirimkan data, tapi tidak ada yang meng-consume

### Unbuffered Channel
Channel yang menampung data max 1 saja. Artinya, ketika data pada channel tersebut telah diisi, namun kita mencoba mengisi kembali channel tersebut, maka akan terjadi blocking.

contoh dapat dilihat di file `channel.go` dan di `channel-fromnet.go`

### Buffered Channel

Buffered channel merupakan cara agar membatasi jumlah data yang masuk pada channel. Caranya adalah dengan menambahkan jumlah angka di bagian param terakhir ketika membuat channel. Ketika channel sudah mencapai limit buffer, barulah goroutine mulai mengeksekusi program.

Pada buffered channel, blocking akan terjadi bila kita mengirim data ke channel. Untuk mencegah ini, consume data dari channel terlebih dahulu, baru kirim kembali data ke channel.

```
var channel = make(chan int, 10)
```
angka `10` di atas menandakan bahwa terdapat **buffer sebesar 11**.

