// Nama: Septiandi Nugraha
// Asal Kampus: Institut Teknologi Telkom Purwokerto

// soal nomor 2
// Buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka, seperti dibawah ini!

package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&angka)

	fmt.Printf("Faktor dari %d adalah: ", angka)
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}