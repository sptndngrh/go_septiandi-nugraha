// Soal 1
// Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri.
// 2 dan 3 adalah bilangan prima. 4 bukan bilangan prima karena 4 bisa dibagi 2. Kamu diminta untuk membuat fungsi untuk menentukan bahwa sebuah bilangan termasuk bilangan prima atau tidak.
// Buatlah solusi yang lebih optimal, dengan kompleksitas lebih cepat dari O(n) / O(n/2).

package main

import "fmt"

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	i := 5
	for i*i <= number {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	var input int
	fmt.Print("Input: ")
	fmt.Scan(&input)

	if primeNumber(input) {
		fmt.Println("Output: Bilangan Prima")
	} else {
		fmt.Println("Output: Bukan Bilangan Prima")
	}

	// Sample case test
	fmt.Println("===========================")
	fmt.Println("Sample case test") 
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13)) // true
	fmt.Println(primeNumber(17)) // true
	fmt.Println(primeNumber(20)) // false
	fmt.Println(primeNumber(35)) // false
}
