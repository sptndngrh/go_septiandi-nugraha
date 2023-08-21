// Nama: Septiandi Nugraha
// Asal Kampus: Institut Teknologi Telkom Purwokerto

// soal nomor 1
// Buatlah sebuah program untuk mencetak segitiga asterik seperti dibawah ini!

package main

import "fmt"

func main() {
	input := 5

	for a := 1; a <= input; a++ {
		for d := 1; d <= input-a; d++ {
			fmt.Print(" ")
		}
		for e := 1; e <= 1*a; e++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}