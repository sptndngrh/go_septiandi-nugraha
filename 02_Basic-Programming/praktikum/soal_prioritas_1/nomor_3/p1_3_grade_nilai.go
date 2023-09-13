// Nama: Septiandi Nugraha
// Asal Kampus: Institut Teknologi Telkom Purwokerto

// soal nomor 3
// Buatlah sebuah program untuk menentukan grade dari sebuah nilai, dengan ketentuan sebagai berikut:
// Nilai 80 - 100: A
// Nilai 65 - 79: B
// Nilai 50 - 64: C
// Nilai 35 - 49: D
// Nilai 0 - 34: E
// Nilai kurang dari 0 atau lebih dari 100 maka tampilkan 'Nilai Invalid'

package main

import (
	"fmt"
)

func main() {
	var nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)

	var peringkat string

	switch {
	case nilai >= 80 && nilai <= 100:
		peringkat = "A"
	case nilai >= 65 && nilai <= 79:
		peringkat = "B"
	case nilai >= 50 && nilai <= 64:
		peringkat = "C"
	case nilai >= 35 && nilai <= 49:
		peringkat = "D"
	case nilai >= 0 && nilai <= 34:
		peringkat = "E"
	default:
		peringkat = "Nilai Invalid"
	}

	fmt.Printf("Grade: %s\n", peringkat)
}