// Soal 1
// Diberi bilangan bulat n, kembalikan array ans dengan panjang n + 1 sehingga untuk setiap i (0 <= i <= n),
// ans[i] adalah bilangan 1 dalam representasi biner dari i
// Input: n = 2
// Output: [0,1,10]
// Input: n = 3
// Output: [0,1,10, 11]

package main

import (
	"fmt"
)

func bin_digits(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		for i := 0; i <= n; i++ {
			if i == 0 {
				ans[i] = "0"
			} else {
				bin := ""
				temp := i
				for temp > 0 {
					remind := temp % 2
					bin = fmt.Sprintf("%d%s", remind, bin)
					temp /= 2
				}
				ans[i] = bin
			}
		}
	}
	return ans
}


func main() {
	n := 2
	ans := bin_digits(n)
	fmt.Println(ans)

	n = 3
	ans = bin_digits(n)
	fmt.Println(ans)
}
