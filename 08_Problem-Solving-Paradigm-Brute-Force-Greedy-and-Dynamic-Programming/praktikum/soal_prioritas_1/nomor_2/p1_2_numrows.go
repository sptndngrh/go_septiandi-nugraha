// Soal 2
// Diberikan bilangan bulat numRows, kembalikan numRows pertama dari segitiga Pascal.
// Dalam segitiga Pascal, setiap angka adalah jumlah dari dua angka tepat di atasnya seperti yang ditunjukkan pada gambar berikut:
// Input: numRows = 5
// Output: [[1], [1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

package main

import (
	"fmt"
)

func segitigaPascal(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	segitiga := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		segitiga = append(segitiga, []int{1})
		for j := 1; j < i; j++ {
			segitiga[i] = append(segitiga[i], segitiga[i-1][j-1] + segitiga[i-1][j])
		}
		segitiga[i] = append(segitiga[i], 1)
	}
	return segitiga
}

func main() {
	segitiga := segitigaPascal(5)
	fmt.Println(segitiga)
}