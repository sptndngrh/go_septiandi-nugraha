// Soal 4
// Buatlah program untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan.
// Gunakan fungsi dengan multiple return, pointer untuk referencing maupun deferencing!

package main

import "fmt"


func getMinMax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}
	return min, max
}


func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	fmt.Println()
	
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai Max : ", max)
	fmt.Println("Nilai Min : ", min)
	}