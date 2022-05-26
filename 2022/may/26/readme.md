## Proses, Thread, dan Green-Thread

**1. Process/Proses**

Proses merupakan suatu operasi yang dijalankan dengan menggunakan alokasi memori yang berbeda. Sehingga, jika terjadi *hang* atau error pada 1 proses, hal tersebut tidak akan mengganggu proses yang lain.

**contoh analogi:**
terdapat 3 orang ingin menggambar secara bersamaan. Maka, kita bisa memberikan masing-masing orang selembar kertas dan pensil untuk menggambar.

kekurangan dari proses adalah, banyak makan memori dan membutuhkan waktu lebih lama untuk membuat suatu proses karena masih perlu mengalokasikan memori dulu sebelum memulai proses.

**2. Thread**

Thread merupakan suatu operasi yang berjalan dengan menggunakan alokasi memori yang sama antara satu thread dengan thread lainnya. Sehingga, thread dapat lebih menghemat memori dan lebih cepat karena tidak memerlukan waktu untuk alokasi memori baru ketika membuat suatu thread baru.

**contoh analogi:**
terdapat 3 orang ingin menggambar secara bersamaan. Maka, kita bisa menyiapkan 1 lembar kertas, dan 3 pena. Masing-masing orang tersebut dapat menggambar pada 1 lembar kertas secara bersamaan.

kekurangannya adalah, karena thread mengakses memori yang sama, diperlukan komunikasi antar thread agar tidak saling tumpang tindih.

**3. Green Thread**

Untuk memahami hal ini, mari pahami beberapa hal tentang bagaimana sistem operasi bekerja.

misalkan terdapat 4 jobs dalam 1 queue of jobs, sistem operasi bekerja dengan memilih job mana yang mau dikerjakan terlebih dahulu di prosesor. Ketika sampai di titik dimana job tersebut perlu dihentikan sementara, sistem operasi akan memasukkan ulang job tersebut ke dalam antrian job dan mengerjakan job yang lain.

kegiatan mengganti job satu ke job lain ini disebut dengan **context switch**, dan lama waktu yang diperlukan untuk mengganti job disebut dengan **context switch overhead**.

jika job yang dikerjakan ada banyak, maka hal ini akan mengakibatkan *context switch overhead* menjadi besar. Sehingga, kecepatan untuk menyelesaikan suatu job akan turun.

**Green Thread** berusaha mengatasi hal ini dengan cara mengurangi *context switch overhead*.

dalam pembahasan thread, terdapat 2 jenis thread.
- Kernel Level Thread, thread yang diproses di level Sistem Operasi (context switch dipilih oleh OS seperti contoh di atas)
- User Level Thread, thread yang diproses di level user (context switch diatur oleh program). Konsep inilah yang digunakan pada green thread.

Pada Green Thread, kita memprogram thread kita agar dapat bekerja secara concurrent dan menentukan kapan kita perlu *switching context*.

Sehingga, yang terjadi adalah:

beberapa user-level thread yang sudah diatur *switching context*-nya masuk ke 1 kernel-level thread -> kernel-level thread diproses CPU -> tidak ada context switching yang terjadi pada kernel-level karena sudah diatur oleh green-thread -> proses selesai.

Kelemahan dari green thread adalah, Sistem Operasi tidak tahu kapan harus *context switching*. Sehingga, jika ada kesalahan dalam switching context pada user-level thread tadi, akan menyebabkan program berjalan dengan tidak semestinya.

## Thread pada GoLang
Golang menganut hybrid system thread dimana golang menyiapkan kernel-level thread pada tiap core CPU komputer kita, masing-masing kernel-level thread memiliki beberapa green-thread. Jadi, ketika ada proses yang sedang berjalan dan proses tersebut harus menunggu user input atau yang lainnya, kernel-level thread dari proses tersebut akan keluar dari antrian, lalu golang akan membuat kernel-level thread baru dimana pada thread baru ini, green-thread dari thread yang dikeluarkan dari antrian tadi, yang tidak memproses proses diatas, akan dipindahkan ke kernel-level thread yang baru saja dibuat.

**rangkaiannya:**
1. Kernel-level thread siap pada masing-masing core CPU, masing2 thread terdapat green-thread.
2. jika ada pekerjaan, pekerjaan tersebut masuk ke green-thread.
3. ketika pekerjaan tersebut harus menunggu, thread akan dipindahkan dari antrian dan golang akan membuat thread baru. Lalu, green-thread yang tidak terpakai/melakukan proses lain selain pekerjaan yang menunggu di atas, akan dipindahkan ke thread yang baru.