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

Sebuah notasi untuk menghitung kompleksitas dari suatu kode. Misal dalam 1 fungsi, terdapat 3 operasi dengan input `array a`:

> 1 + a[0] -> O(1)
> 
> sum(a) -> O(n)
> 
> pair(a) -> O(n<sup>2</sup>)

Maka, penulisan big O notationnya adalah:

> **O(n<sup>2</sup> + n + 1)**

**TAPI**, dalam konsep asimptotik, Jika ada konstanta atau 2 variable yang sama dalam 1 notasi, kita hanya mencatat 1 variable terbesar saja. Sehingga, notasi di atas menjadi:

> **O(n<sup>2</sup>)**

### Catatan
> Operasi yang tidak berpengaruh terhadap perubahan data input, tidak menambah angka notasi. Misal terdapat 3 operasi dengan input array a:
> mendeklarasikan variable
> membuat konstanta
> melakukan a[0] + 1
> 
> Maka kompleksitas di atas bukanlah O(3), tapi tetap O(1) karena baik berapapun input a yang diberikan, kita tetap membuat variable, tetap membuat konstanta, dan melakukan 1 proses saja yaitu a[0] + 1. **Tidak ada yang namanya kompleksitas O(2), O(3), dst, Karena hal tersebut konstan. Sehingga, dianggap kompleksitas tersebut sama dengan O(1)**.

Jika ada 2 input pada 1 fungsi, misal input array n dan m, maka notasinya ditulis berdasarkan variable terbesar dari masing-masing input tadi. Misal terdapat notasi:

> **O(m<sup>2</sup>+m+2n+3)**

Maka, kompleksitas di atas harus diubah menjadi:

> **O(m<sup>2</sup>+n)**