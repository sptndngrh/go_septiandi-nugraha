// soal nomor 1
// Buatlah sebuah program untuk memeriksa apakah sebuah kata adalah palindrome atau bukan, serta coba terapkan scanner untuk menangkap inputan dari console, seperti dibawah ini!

package main

import (
	"fmt"
	"strings"
)


func main() {
	fmt.Print("Program pengecekan kata Palindrom \n")
	var kata string
	fmt.Print("masukkan sebuah kata: ")
	fmt.Scan(&kata)

	if palindrome(kata) {
		fmt.Printf("%s merupakan palindrome.\n", kata)
	} else {
		fmt.Printf("%s tidak termasuk palindrome.\n", kata)
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
