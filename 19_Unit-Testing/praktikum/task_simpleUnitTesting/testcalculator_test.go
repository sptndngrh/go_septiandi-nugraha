package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	c := calculator{}

	// pengetesan untuk penjumlahan
	result := c.Addition(5, 3)
	if result != 8 {
		t.Errorf("Hasilnya yaitu: 8, Hasil Run :  %f", result)
	}

	// pengetesan untuk pengurangan
	result = c.Subtraction(5, 3)
	if result != 2 {
		t.Errorf("Hasilnya yaitu: 2, Hasil Run :  %f", result)
	}

	// pengetesan untuk pembagian
	result, err := c.Division(6, 2)
	if err != nil || result != 3 {
		t.Errorf("Hasilnya yaitu: 3 tanpa error, tetapi hasilnya %f dengan error: %v", result, err)
	}

	// pengetesan untuk pembagian oleh 0 (error case)
	_, err = c.Division(6, 0)
	if err == nil {
		t.Errorf("Hasilnya yaitu ada error saat membagi oleh 0, tetapi tidak ada error yang ditemukan")
	}

	// pengetesan untuk perkalian
	result = c.Multiplication(5, 3)
	if result != 15 {
		t.Errorf("Hasilnya yaitu: 15, Hasil Run :  %f", result)
	}

	// Tambahkan pengetesan untuk mencakup jalur alternatif dalam Division
	result, err = c.Division(3, 0)
	if err == nil {
		t.Errorf("Hasilnya yaitu ada error saat membagi oleh 0, tetapi tidak ada error yang ditemukan")
	}
}