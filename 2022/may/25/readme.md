# parallelism, thread dan concurrent dalam programming


**Pertanyaan:**
1. Apa itu parallelism dalam programming?
2. Apa itu thread dalam programming?
3. Apa itu concurrent dalam programming?


**Jawaban:**

1. Parallelism adalah operasi yang berjalan bersamaan dalam satu waktu pada beberapa core CPU. Hal ini tidak bisa terjadi dalam single-core CPU karena 1 core hanya dapat memproses 1 proses.

2. Thread merupakan bagian dari suatu operasi atau dapat disebut sub-operasi. Threading merupakan bagaimana beberapa thread dapat berjalan dalam 1 core CPU.
3. Concurrent/concurrency adalah bagaimana beberapa operasi dapat berjalan bersamaan dalam suatu CPU, baik single-core maupun multiple-core CPU. 
   1. Dalam single core CPU, concurrency dapat terjadi dengan menghentikan sementara suatu proses, dan mengerjakan proses lainnya, kemudian melanjutkan proses yang sebelumnya tertunda.
   2. Dalam multiple core CPU (dual-core, quad-core, dll.), concurrency dapat terjadi dengan menjalankan suatu operasi bersamaan dalam masing-masing core CPU.