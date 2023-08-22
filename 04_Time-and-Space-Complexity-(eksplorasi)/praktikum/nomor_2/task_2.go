// Soal 2
// Terdapat dua bilangan integer yaitu x dan n. Buatlah sebuah fungsi untuk melakukan perhitungan perpangkatan (x^n dibaca x pangkat n).
// Time complexity dari sebuah fungsi perpangkatan harus lebih cepat dari O(n). Contoh time complexity yang optimal adalah logaritmik.

package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	temp := pow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	} else {
		return x * temp * temp
	}
}

func main() {
	var x, n int
	fmt.Print("Input x = ")
	fmt.Scan(&x)
	fmt.Print("Input n = ")
	fmt.Scan(&n)

	result := pow (x, n)
	fmt.Printf("Output: %d\n", result)

	fmt.Println("===========================")
	fmt.Println("Sample case test") 
	fmt.Println(pow(2, 3)) // 8
	fmt.Println(pow(5, 3)) // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5)) // 32
	fmt.Println(pow(7, 3)) // 343
}