# Struktur Data

Struktur data adalah cara mengatur kumpulan data yang dapat dimanipulasi dengan suatu atau beberapa fungsi/operasi.

## Kompleksitas Data
- Analisa kompleksitas
  
Dilakukan untuk menganalisa apakah suatu algoritma itu baik atau tidak dan membandingkannya dengan algoritma yang lain. Biasanya menggunakan **time complexity** dan **space complexity**,

- time complexity = seberapa cepat suatu algoritma dapat menyelesaikan masalah.
- space complexity = seberapa banyak memori yang digunakan oleh suatu algoritma dalam menyelesaikan masalah.

Kedua hal diatas biasa dituliskan dalam big O notation.

## Memori

- Memori merupakan tempat dimana suatu data tersimpan.
- Memori memiliki segmen-segmen terpisah yang disebut slot memori. 1 slot memori dapat menampung 8 bit (binary digit) atau 1 byte. Contoh 8 bit: 0000 0001
- Jika suatu data tersimpan dalam slot memori, dengan ukuran lebih dari 1 byte, maka data tersebut akan tersimpan ke lebih dari 1 slot memori yang berurutan. Misal: 32bit akan tersimpan di 4 slot yang berurutan.
- slot memori juga dapat menyimpan alamat dari slot memori lainnya (pointer).

## Big O Notation

Sebuah notasi untuk mengetahui kompleksitas dari suatu kode. Misal dalam 1 fungsi, terdapat 3 operasi dengan input `array a`:

> 1 + a[0] -> O(1)
> 
> sum(a) -> O(n)
> 
> pair(a) -> O(n<sup>2</sup>)

Maka, penulisan big O notationnya adalah:

> **O(n<sup>2</sup> + n + 1)**

**TAPI**, dalam konsep asimptotik, Jika ada konstanta atau 2 variable yang sama dalam 1 notasi, kita hanya mencatat 1 variable terbesar saja. Sehingga, notasi di atas menjadi:

> **O(n<sup>2</sup>)**

### **Catatan**
> Operasi yang tidak berpengaruh terhadap perubahan data input, tidak menambah angka notasi. Misal terdapat 3 operasi dengan input array a:
> - mendeklarasikan variable
> - membuat konstanta
> - melakukan a[0] + 1
> - melakukan a[0] - 1
> - melakukan a[1] - a[2]
> Maka kompleksitas di atas bukanlah O(5), tapi tetap O(1) karena baik berapapun input a yang diberikan, kita tetap melakukan jumlah proses yang sama. **Tidak ada yang namanya kompleksitas O(2), O(3), dst, Karena hal tersebut konstan. Sehingga, dianggap kompleksitas tersebut sama dengan O(1)**.

Jika ada 2 input pada 1 fungsi, misal input array n dan m, maka notasinya ditulis berdasarkan variable terbesar dari masing-masing input tadi. Misal terdapat notasi:

> **O(m<sup>2</sup>+m+2n+3)**

Maka, kompleksitas di atas harus diubah menjadi:

> **O(m<sup>2</sup>+n)**

Dari terbaik ke terburuk, kompleksitas suatu algoritma dapat diurutkan menjadi:
- O(1) -> konstan. Tidak terpengaruh ke perubahan input.
- O(log(n))
- O(n) -> linear. 
- O(n log(n))
- O(n<sup>2</sup>), O(n<sup>3</sup>), dst.
- O(2<sup>n</sup>)
- O(n!) -> faktorial

## Log

Dalam computer science, `log N` berarti **log<sub>2</sub> N** (menggunakan basis 2, bukan 10). Contoh:

Jika terdapat program dengan kompleksitas `log N` dan menerima input panjang array 8. Maka jumlah operasi yang akan dijalankan program tersebut adalah:

> log 8 dapat diartikan menjadi 2<sup>?</sup> = 8
> 
> ? = 3

Dari perhitungan di atas, kita bisa mengetahui bahwa program tersebut melakukan 3 operasi ketika menerima input array sepanjang 8.

Ketika panjang array tersebut digandakan, kita hanya perlu menambahkan `? + 1`. Sehingga, jika panjang array saat ini 16, maka program tersebut melakukan 4 operasi.

Itulah kenapa `O(log(N))` lebih baik dibanding linear.

