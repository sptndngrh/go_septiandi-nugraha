// Soal 1
// Tulis program untuk mengkonversi dari angka normal ke Angka Romawi!
// Input: 4
// Output: IV

// Input: 9
// Output: IX

// Input: 23
// Output: XXIII

// Input: 2021
// Output: MMXXI

// Input: 1646
// Output: MDCXLVI

package main

import "fmt"

func intToRoman(num int) string {
	romanNumerals := []struct {
		value  int
		symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"},
		{1, "I"},
	}

	romanNum := ""
	for _, pair := range romanNumerals {
		for num >= pair.value {
			romanNum += pair.symbol
			num -= pair.value
		}
		if num == 0 {
			break
		}
	}
	return romanNum
}

func main() {
	fmt.Println(intToRoman(4))    // Output: IV
	fmt.Println(intToRoman(9))    // Output: IX
	fmt.Println(intToRoman(23))   // Output: XXIII
	fmt.Println(intToRoman(2021)) // Output: MMXXI
	fmt.Println(intToRoman(1646)) // Output: MDCXLVI
}
