# Materi 17 QRM and Code Structure
Tiga poin penting terkait dengan ORM, khususnya dengan GORM dalam bahasa pemrograman Go (Golang), adalah:

> ORM (Object-Relational Mapping):
ORM adalah teknik pemrograman yang memungkinkan konversi data antara sistem tipe yang tidak kompatibel menggunakan bahasa pemrograman berbasis objek. GORM adalah salah satu library ORM yang digunakan dalam Go untuk mempermudah interaksi antara aplikasi Go dan database relasional.

>Keunggulan ORM:
- Mengurangi jumlah query yang berulang-ulang.
- Mengambil data secara otomatis ke dalam objek yang siap digunakan.
- Memiliki cara sederhana untuk menyaring data sebelum disimpan di database.
- Beberapa ORM memiliki fitur permintaan cache.

> Kekurangan ORM:
- Menambahkan lapisan kode tambahan dan biaya proses overhead.
- Memuat data hubungan yang mungkin tidak diperlukan.
- Kueri SQL yang kompleks dapat memakan waktu lama untuk ditulis dengan ORM, terutama jika melibatkan banyak tabel join.
- Beberapa fungsi SQL khusus yang terkait dengan satu vendor database mungkin tidak didukung, atau tidak ada fungsi tertentu yang tersedia (misalnya, MySQL: Force Index).

Selain itu, dalam konteks GORM, beberapa konsep penting meliputi "Database Migration," yang digunakan untuk memperbarui struktur database agar sesuai dengan perubahan versi aplikasi, dan "DB Relation," yang mencakup jenis-jenis hubungan antar tabel dalam GORM seperti "Belongs to," "Has one," "Has many," dan "Many to many."