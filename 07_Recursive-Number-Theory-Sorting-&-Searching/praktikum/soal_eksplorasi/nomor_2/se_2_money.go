// Soal 2
// Kamu diminta untuk membeli sebuah barang, dan tantangan kali ini kamu harus bisa membeli barang dengan jumlah maksimum dengan uang yang kamu miliki.
// Program ini menerima money sebagai parameter pertama dan yang kedua berupa list harga produk yang dapat dibeli. Kamus harus menampilkan nilai jumlah barang yang bisa dibeli.
// Sample Test Case
// Input: money = 50000, productPrice = [25000, 25000, 10000, 14000]
// Output: 3
// Input: money = 50000, productPrice = 15000, 10000, 12000, 5000, 3000]
// Output 5

package main

import "fmt"


func MaximumBuyProduct(money int, productPrice []int) int {
	swap := func(arr []int, i, j int) {
        arr[i], arr[j] = arr[j], arr[i]
    }

	selectionSort := func(arr []int) {
        n := len(arr)
        for i := 0; i < n-1; i++ {
            minIndex := i
            for j := i + 1; j < n; j++ {
                if arr[j] < arr[minIndex] {
                    minIndex = j
                }
            }
            if minIndex != i {
                swap(arr, i, minIndex)
            }
        }
    }

	selectionSort(productPrice) 

	count := 0
	totalPrice := 0
	for _, price := range productPrice {
		if totalPrice+price <= money {
			totalPrice += price
			count++
		} else {
			break
		}
	}

	return count
}


func main() {
	fmt.Println(MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})) // 3
	fmt.Println(MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})) // 4
	fmt.Println(MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})) // 4
	fmt.Println(MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})) // 1
	fmt.Println(MaximumBuyProduct(0, []int{10000, 30000})) // 0
}

