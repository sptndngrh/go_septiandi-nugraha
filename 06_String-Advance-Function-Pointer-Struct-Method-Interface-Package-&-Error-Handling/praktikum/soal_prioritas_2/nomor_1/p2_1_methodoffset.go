// Soal 1
// Buatlah sebuah method dengan parameter offset bertipe integer dan input bertipe string. Method ini menghasilkan
// sebuah string baru yang dimana setiap huruf dilakukan pergeseran berdasarkan offset yang merupakan jumlah pergeseran hurufnya.
// String input diasumsikan berisi huruf kecil saja dan spasi. Sebagai contoh, ketika terdapat huruf z yang digeser dengan offset sebesar 3 maka huruf hasil pergeseran adalah c.

// Daftar karakter ASCII dapat dilihat di link berikut.

// Berdasarkan referensi ASCII, huruf a memiliki kode 97, huruf b memiliki kode 98, z memiliki kode 122. 
// Anda bisa menggunakan operator modulo jika diperlukan.

package main

import "fmt"


func caesar(offset int, input string) string {
	result := ""

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shifted := 'a' + (char-'a'+rune(offset))%26
			result += string(shifted)
		} else if char >= 'A' && char <= 'Z' {
			shifted := 'A' + (char- 'A' +rune(offset))%26
			result += string(shifted)
		} else {
			result += string(char)
		}
	}
	return result
}


func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}


