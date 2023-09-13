// Nama: Septiandi Nugraha
// Asal Kampus: Institut Teknologi Telkom Purwokerto

// soal nomor 1
// Buatlah sebuah program untuk memeriksa apakah sebuah kata adalah palindrome atau bukan, serta coba terapkan scanner untuk menangkap inputan dari console, seperti dibawah ini!

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main() {
	fmt.Println("Apakah Palindrome?")
	fmt.Print("Masukkan kata: ")

	// Membuat objek Scanner untuk membaca seluruh baris input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kata := scanner.Text()

	if palindrome(kata) {
		fmt.Print("Captured: ", kata)
		fmt.Printf("\nPalindrome")
	} else {
		fmt.Print("Captured: ", kata)
		fmt.Printf("\nBukan merupakan palindrome")
	}
}

func palindrome(str string) bool {
	str = strings.ToLower(str)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
