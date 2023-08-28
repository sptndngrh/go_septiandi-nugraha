// Soal 3
// Buatlah program untuk menentukan substring yang sama di antara keduanya, dengan diberikan dua string A dan B!

package main


import (
	"fmt"
)


func Compare(a, b string) string {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	commonSubstring := ""

	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			break
		}
		commonSubstring += string(a[i])
	}
	return commonSubstring
}


func main() {
	fmt.Println(Compare("AKA", "AKASHI")) //AKA
	fmt.Println(Compare("KANGOORO", "KANG")) //KANG
	fmt.Println(Compare("KI", "KIJANG")) //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA")) //ILA
}