// Soal 4
// Buatlah program yang yang menerapkan mutex! Jenis program yang dibuat bebas (contoh: perhitungan faktorial).

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Harga per liter untuk berbagai produk bensin
	hargaPertalite := 10000.0
	hargaPertamax := 13300.0
	hargaPertamaxGreen95 := 13500.0
	hargaPertamaxTurbo92 := 14000.0

	// Input jumlah uang yang dimiliki oleh pengguna
	var uang float64
	fmt.Print("Masukkan jumlah uang Anda: ")
	_, err := fmt.Scan(&uang)
	if err != nil || uang <= 0 {
		fmt.Println("Masukkan jumlah uang yang valid!")
		return
	}

	// Menggunakan mutex untuk melindungi akses ke data bersama
	var mu sync.Mutex

	// Menggunakan channel dengan buffer untuk menghitung jumlah liter
	ch := make(chan struct {
		produk       string
		literBensin  float64
	}, 6) // Buffer 6 untuk menampung hasil dari 6 goroutine

	// Goroutine untuk menghitung jumlah liter masing-masing produk bensin
	produkBensin := []struct {
		nama  string
		harga float64
	}{
		{"Pertalite (Harga: Rp 10.000)", hargaPertalite},
		{"Pertamax (Harga: Rp 13.300)", hargaPertamax},
		{"Pertamax Green 95 (Harga: Rp 13.500)", hargaPertamaxGreen95},
		{"Pertamax Turbo 92 (Harga: Rp 14.000)", hargaPertamaxTurbo92},
	}

	fmt.Printf("Dengan uang Rp %.0f, Anda dapat membeli:\n", uang)

	var wg sync.WaitGroup
	for _, p := range produkBensin {
		wg.Add(1)
		go func(nama string, harga float64) {
			defer wg.Done()
			mu.Lock()
			liter := uang / harga
			mu.Unlock()
			time.Sleep(1 * time.Second) // Simulasi penundaan penghitungan liter
			ch <- struct {
				produk      string
				literBensin float64
			}{nama, liter}
		}(p.nama, p.harga)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Menampilkan hasil penghitungan liter
	for result := range ch {
		fmt.Printf("%s: %.2f liter\n", result.produk, result.literBensin)
	}

	// Menentukan produk bensin yang paling hemat
	var produkRekomendasi string
	var jumlahLiterTerbanyak float64

	for _, p := range produkBensin {
		liter := uang / p.harga
		if liter > jumlahLiterTerbanyak {
			produkRekomendasi = p.nama
			jumlahLiterTerbanyak = liter
		}
	}

	fmt.Printf("Produk yang hemat dan direkomendasikan adalah %s dengan %.2f liter.\n", produkRekomendasi, jumlahLiterTerbanyak)
}