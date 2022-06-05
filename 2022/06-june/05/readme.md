## Conditional Variable

Merupakan bagian dari library bawaan golang di `sync.newCondition` dimana fungsi ini berguna untuk mengirimkan sinyal ketika fungsi di dalam mutex memiliki kondisi yang tidak memenuhi syarat untuk dilakukannya suatu proses, namun ingin proses tersebut berjalan setelah kondisinya terpenuhi.


Contoh:
Kita memiliki 0 saldo dan melakukan 3 proses secara berurutan:
- menambah saldo 10
- mengurangi saldo 20
- menambah saldo 10

Pada kasus diatas, ketika mengurangi saldo sebanyak 20, akan muncul kesalahan logic dimana kita tidak boleh mengambil uang lebih banyak dari sisa saldo. Sehingga, proses kedua ini tidak mungkin dilakukan. Namun, kita ingin proses ini **SEGERA DILAKUKAN KETIKA SALDO SUDAH CUKUP**.

ketika proses ketiga masuk, saldo saat ini adalah 20. Sehingga, proses kedua tadi sudah bisa dilakukan. Di sinilah kita bisa menggunakan condition variable.

## Cara Kerja
1. Ketika kondisi tidak memenuhi, condition variable akan meng-unlock variable yang terkunci dan akan blocking prosesnya saat ini.
2. Ketika kita sudah melakukan proses yang kemungkinan membuat condition variable terpenuhi, kita mengirimkan signal ke proses no 1.
3. Proses no 1 akan melanjutkan proses.