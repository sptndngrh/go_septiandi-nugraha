package main

import (
	"errors"
)

// Struktur untuk menyimpan fungsi-fungsi kalkulasi
type calculator struct{}

// Penjumlahan
func (c calculator) Addition(a, b float64) float64 {
	return a + b
}

// Pengurangan
func (c calculator) Subtraction(a, b float64) float64 {
	return a - b
}

// Pembagian
func (c calculator) Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Pembagian oleh nol tidak diizinkan")
	}
	return a / b, nil
}

// Pengkalian
func (c calculator) Multiplication(a, b float64) float64 {
	return a * b
}