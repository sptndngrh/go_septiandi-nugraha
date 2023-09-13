// Soal 1
// Buatlah sebuah program menggabungkan 2 array yang diberikan, dan jangan sampai
// terdapat nama yang sama di data yang sudah tergabung tadi!

package main

import (
	"fmt"
	"strings"
)


func ArrayMerge(arrayA, arrayB []string) string {
	mergedMap := make(map[string]bool) 
	merged := []string{}  
	             
	for _, name := range arrayA {
		mergedMap[name] = true
		merged = append(merged, name)
	}

	for _, name := range arrayB {
		if !mergedMap[name] {
			mergedMap[name] = true
			merged = append(merged, name)
		}
	}

	mergedQuote := make([]string, len(merged))
	for i, name := range merged {
		mergedQuote[i] = `"` + name + `"`
	}
	mergedString := "[" + strings.Join(mergedQuote, ", ") + "]"
	return mergedString
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}