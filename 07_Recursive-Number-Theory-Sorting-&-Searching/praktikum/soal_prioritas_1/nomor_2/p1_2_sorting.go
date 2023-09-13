// Soal 2
// Buatlah program yang dapat mengurutkan item berdasarkan jumlah kemunculannya.
// Jika ada item yang duplicate kamu hanya perlu memunculkan sekali, namun kamu perlu menampilkan total kemunculan item tersebut!

package main

import (
	"fmt"
)


type pair struct {
	name string
	count int
}


func MostAppearItem(items []string) []pair {
	freq := make(map[string]int)
	for _, item := range items {
		freq[item]++
	}

	var pairs []pair
	for name, count := range freq {
		pairs = append(pairs, pair{name, count})
	}

	n := len(pairs)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if pairs[j].count < pairs[minIndex].count || (pairs[j].count == pairs[minIndex].count && pairs[j].name > pairs[minIndex].name) {
				minIndex = j
			}
		}
		pairs[i], pairs[minIndex] = pairs[minIndex], pairs[i]
	}

	return pairs
}


func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}

fmt.Println()
	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}

fmt.Println()
pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
for _, list := range pairs {
	fmt.Print(list.name, " -> ", list.count, " ")
	}
}