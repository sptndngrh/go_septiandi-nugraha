
3 point utama dari Clean dan Hexagonal Architecture dalam Golang:

> Pemisahan masalah: Aplikasi dibagi menjadi lapisan yang berbeda, dengan masing-masing lapisan bertanggung jawab untuk tugas tertentu.
> Inversi dependensi: Komponen tidak bergantung pada implementasi spesifik dari komponen lain. Sebaliknya, mereka bergantung pada antarmuka yang didefinisikan oleh komponen lain.
> Pengujian unit: Kode diuji secara menyeluruh menggunakan pengujian unit.

Berikut adalah penjelasan singkat dari setiap poin:

> Pemisahan masalah membantu membuat aplikasi lebih mudah dipelihara dan diuji. Dengan memisahkan masalah menjadi lapisan yang berbeda, yang dapat dengan mudah memahami bagaimana aplikasi bekerja dan memperbaikinya jika terjadi bug.
? Inversi dependensi membantu membuat aplikasi lebih modular dan dapat diuji. Dengan tidak bergantung pada implementasi spesifik dari komponen lain, yang dapat dengan mudah mengganti implementasi tersebut tanpa mempengaruhi komponen lain.
> Pengujian unit membantu memastikan bahwa kode berfungsi dengan benar. Dengan menulis pengujian unit dapat memastikan bahwa kode memenuhi semua persyaratannya.

Berikut adalah beberapa manfaat dari menerapkan Clean dan Hexagonal Architecture dalam Golang:
> Aplikasi yang lebih mudah dipelihara (Maintenance)
> Aplikasi yang lebih dapat diuji
> Aplikasi yang lebih modular
> Aplikasi yang lebih dapat diskalakan
