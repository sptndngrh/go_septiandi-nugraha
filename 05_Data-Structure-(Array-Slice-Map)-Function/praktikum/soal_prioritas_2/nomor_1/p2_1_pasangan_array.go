// Soal 1
// Diberikan array angka yang diurutkan dan jumlah target, temukan pasangan dalam array yang jumlahnya sama dengan target yang diberikan.
// Tulis fungsi untuk mengembalikan indeks dari dua angka (yaitu pasangan) sehingga jika value pada index tersebut dijumlahkan sesuai target yang diberikan.

package main

import "fmt"


func PairSum(arr []int, targeter int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == targeter {
			return []int{left, right}
		} else if sum < targeter {
			left++
		} else {
			right--
		}
	}
	return []int{}
}


func main() {
	// Test cases
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]
}