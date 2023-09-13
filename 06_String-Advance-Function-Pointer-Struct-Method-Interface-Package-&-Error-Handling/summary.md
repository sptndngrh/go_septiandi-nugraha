# Materi String Advance Function Pointer Struct Method Interface Package & Error Handling

> String
String adalah tipe data yang digunakan untuk merepresentasikan urutan karakter. String dalam Go direpresentasikan sebagai serangkaian karakter Unicode. String dapat dideklarasikan dengan menggunakan tanda kutip ganda atau tanda kutip tunggal. Panjang sebuah string dapat dihitung menggunakan fungsi len().
Karakter dalam sebuah string dapat diakses menggunakan indeks. String bersifat immutable, artinya setelah dibuat, Anda tidak dapat mengubah karakter individunya.

> Advance Function & Pointer
Variadic function adalah fungsi yang dapat menerima jumlah argumen yang bervariasi. Variadic function diimplementasikan dengan menggunakan tanda elipsis ... sebelum tipe parameter variadic. Fungsi variadic memungkinkan Anda untuk mengoperasikan sejumlah nilai tanpa harus membatasi jumlah argumen secara eksplisit.

Pointer adalah tipe data yang digunakan untuk mengakses dan mengubah nilai dari suatu variabel melalui alamat memori variabel tersebut. Setiap variabel memiliki lokasi memori yang terkait dengannya. Operator & digunakan untuk mendapatkan alamat memori dari suatu variabel. Operator * digunakan untuk mendapatkan nilai dari variabel yang ditunjuk oleh pointer (dereferensi). Pointer sering digunakan dalam fungsi untuk mengubah nilai variabel di luar fungsi tersebut.  

> Package & Error Handling
Package adalah cara untuk mengorganisir dan mengelompokkan kode program di Golang. Setiap file Go harus dimasukkan ke dalam package tertentu. Package membantu mengatur kode program ke dalam unit yang lebih terstruktur dan mudah dikelola. Package juga memungkinkan untuk mengimpor kode dari package lain untuk penggunaan ulang dan modularitas.

Error handling adalah konsep penting dalam pemrograman yang memungkinkan program untuk mengatasi situasi yang tidak terduga atau kesalahan saat program dieksekusi. Dalam Go, error dianggap sebagai tipe data, dan error handling dilakukan dengan mengembalikan nilai error sebagai bagian dari hasil fungsi. Fungsi yang mengembalikan error biasanya memiliki dua nilai kembalian, yaitu hasil operasi dan error. Error handling dilakukan dengan memeriksa nilai error yang dikembalikan oleh fungsi dan mengambil tindakan yang sesuai.