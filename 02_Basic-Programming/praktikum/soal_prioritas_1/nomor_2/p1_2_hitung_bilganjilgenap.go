// Nama: Septiandi Nugraha
// Asal Kampus: Institut Teknologi Telkom Purwokerto

// soal nomor 2
// Buatlah sebuah program untuk menentukan apakah sebuah bilang adalah bilang ganjil atau genap.

package main

import "fmt"

func main() {
	var bil int

	fmt.Print("Masukkan bilangan yang anda kehendaki: ")
	fmt.Scanln(&bil)

	if bil%2 == 0 {
		fmt.Printf("%d merupakan bilangan genap. \n", bil)
	} else {
		fmt.Printf("%d merupakan bilangan ganjil. \n", bil)
	}
}