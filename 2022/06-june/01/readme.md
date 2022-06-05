## Thread Pool

Konsepnya sama dengan worker pool, dimana ada master thread dan ada worker thread. Master thread akan memasukkan pekerjaan ke buffered channel, kemudian para worker yang sudah dibuat sebelumnya akan mengerjakan task berdasarkan data dari channel tersebut.