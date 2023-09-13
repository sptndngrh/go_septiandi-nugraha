// Soal 2
// Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan. Ubahlah penulisan kode berikut menjadi lebih terbaca dan lebih rapi!


package main


import "fmt"


// Definisikan tipe data dari struktur Kendaraan
type Kendaraan struct {
	TotalRoda 		int
	KecepatanPerJam int
}


// Definisikan tipe data mobil dari warisan struktur tipe data kendaraan
type Mobil struct {
	Kendaraan
}


// Metode untuk mobil ketika berjalan dengan menambah kecepatan di tipe data mobil
func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}


// Metode untuk menambah kecepatan mobil dengan jumlah tertentu pada tipe data mobil
func (m *Mobil) TambahKecepatan(KecepatanBaru int) {
	m.KecepatanPerJam += KecepatanBaru
}


func main() {
	// Inisialisasi variabel mobil ketika dalam kecepatan tinggi
	mobilCepat := Mobil{}

	// Ketika mobil cepat berjalan tiga kali
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	// Inisialisasi variabel mobil ketika dalam kecepatan lambat
	mobilLambat := Mobil{}

	// mobil ketika lamban berjalan satu kali
	mobilLambat.Berjalan()

	// Mengeluarkan keterangan kecepatan mobil ketika output program dijalankan
	fmt.Printf("Kecepatan mobil cepat: %d km/jam\n", mobilCepat.KecepatanPerJam)
	fmt.Printf("Kecepatan mobil lambat: %d km/jam\n", mobilLambat.KecepatanPerJam)
}


