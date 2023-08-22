// Soal 3
// Buatlah sebuah program untuk menampilkan
// angka yang hanya muncul satu kali.

package main

import (
	"fmt"
	"strings"
)



func munculSekali(angka string) []int {
	counts := make(map[int]int)
	var result []int

	digits := strings.Split(angka, "")
	for _, digitStr := range digits {
		var digit int
		fmt.Sscanf(digitStr, "%d", &digit)
		counts[digit]++
	}

	for digit, count := range counts {
		if count == 1 {
			result = append(result, digit)
		}
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(munculSekali("1234123")) // [4]
	fmt.Println(munculSekali("76523752")) // [6 3]
	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]
}