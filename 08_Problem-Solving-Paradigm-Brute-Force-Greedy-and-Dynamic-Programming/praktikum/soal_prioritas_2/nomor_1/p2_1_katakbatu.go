// Soal 1
// Ada katak yang awalnya berada di atas Batu 1. Dia akan mengulangi tindakan berikut beberapa kali untuk mencapai batu N. 
// Jika katak sedang berada di Batu i, lompat ke Batu i + 1 atau Batu i + 2. Di sini, biaya | hi - hj | terjadi, di mana j adalah batu untuk mendarat.
// Temukan biaya total minimum yang mungkin dikeluarkan sebelum katak mencapai Batu N.

package main

import (
	"fmt" 
)


func Frog(jumps []int) int {
	n := len(jumps)
	if n <= 2 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	j := make([]int, n)
	j[0] = 0
	j[1] = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}(jumps[1] - jumps[0])

	for i := 2; i < n; i++ {
		j[i] = min(j[i-1]+func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		}(jumps[i]-jumps[i-1]), j[i-2]+func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		}(jumps[i]-jumps[i-2]))
	}

	return j[n-1]
}

func main() {
fmt.Println(Frog([]int{10, 30, 40, 20})) 			// 30
fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) 	// 40
}