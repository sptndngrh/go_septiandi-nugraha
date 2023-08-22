// Soal 2
// Buatlah sebuah program yang dapat menghitung berapa 
// banyak sebuah string yang ada didalam sebuah slice!

package main

import "fmt"


func Mapping(slice []string) map[string]int {
	counts := make(map[string]int)
	for _, s := range slice {
		counts[s]++
	}
	return counts
}


func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"})) // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{})) // map[]
}


