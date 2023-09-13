// Soal 3
// Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri. 
// Angka 2 dan 3 adalah bilangan prima. Angka 4 bukan bilangan prima karena 4 bisa dibagi 2.
// Sepuluh deret bilangan prima yang pertama adalah 2, 3, 5, 7, 11, 13, 17, 19, 23 dan 29.
// Buatlah sebuah fungsi bernama getPrime yang menampilkan bilangan prima sesuai dengan deret urutannya.

package main

import (
	"fmt"
)

func primeX(number int) int {
	count := 0
	currNumber := 2

	for {
		isPrime := true
		for i := 2; i <= currNumber/2; i++ {
			if currNumber%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
		if count == number {
			return currNumber
		}
		currNumber++
	}
}

func main() {
	fmt.Println(primeX(1)) // 2
	fmt.Println(primeX(5)) // 11
	fmt.Println(primeX(8)) // 19
	fmt.Println(primeX(9)) // 23
	fmt.Println(primeX(10)) // 29
}