// Nama: Septiandi Nugraha
// Asal Kampus: Institut Teknologi Telkom Purwokerto

// soal nomor 4
// Buatlah sebuah program yang mencetak angka dari 1 sampai 100 
// dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk 
// kelipatan '5' cetak "Buzz”. Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz

package main

import "fmt"

func main() {
	for angka := 1; angka<= 100; angka++ {
		if angka%3 == 0 && angka%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if angka%3 == 0 {
			fmt.Println("Fizz")
		} else if angka%5 == 0{
			fmt.Println("Buzz")
		} else {
			fmt.Println(angka)
		}
	}
}