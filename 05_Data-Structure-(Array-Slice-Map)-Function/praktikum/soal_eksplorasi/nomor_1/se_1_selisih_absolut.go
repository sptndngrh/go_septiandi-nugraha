// Soal 1
// Diberi matriks persegi, buatlah program untuk menghitung selisih absolut antara jumlah diagonalnya.
// a. Input:
// 	  1 2 3
// 	  4 5 6
// 	  9 8 9
// b. Penjelasan: Diagonal kiri ke kanan : 1+5+9 = 15 .
//    Diagonal kanan ke kiri : 3+5+9 = 17 . Perbedaan mutlak adalah |15 - 17| = 2.

package main

import (
	"fmt"
	"math"
)

func diagonalDiff(matriks [][]int) int {
	n := len(matriks)
	diagonalKiri := 0
	diagonalKanan := 0

	for i := 0; i < n; i++ {
		diagonalKiri += matriks[i][i]
		diagonalKanan += matriks[i][n-i-1]
	}
	return int(math.Abs(float64(diagonalKiri - diagonalKanan)))
}

func main() {
	matriks := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	result := diagonalDiff(matriks)
	fmt.Println("Perbedaan diagonal mutlak adalah", result)
}