// Soal 1
// Hitung frekuensi huruf dalam teks menggunakan perhitungan paralel (Bersamaan).

package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time" 
)

func main() {
	fmt.Print("Masukkan teks: ")
	kalimat := bacaKalimat()

	// Cetak pesan loading sekali sebelum menghitung frekuensi huruf
	fmt.Println("Loading...")

	frekuensi := hitungFrekuensiHuruf(kalimat)

	// Tunggu 1 detik sebelum menampilkan hasil
	time.Sleep(3 * time.Second)

	tampilkanFrekuensi(frekuensi)
}


func bacaKalimat() string {
	reader := bufio.NewReader(os.Stdin)
	kalimat, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return kalimat
}

func hitungFrekuensiHuruf(kalimat string) map[rune]int {
	frekuensi := make(map[rune]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	potongan := bagiKalimat(kalimat, 4)

	for _, pot := range potongan {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()
			perhitunganPotongan(p, &mu, frekuensi)
		}(pot)
	}

	wg.Wait()
	return frekuensi
}

func bagiKalimat(kalimat string, jumlahPotongan int) []string {
	panjangPotongan := len(kalimat) / jumlahPotongan
	potongan := make([]string, 0)

	for i := 0; i < len(kalimat); i += panjangPotongan {
		akhir := i + panjangPotongan
		if akhir > len(kalimat) {
			akhir = len(kalimat)
		}
		potongan = append(potongan, kalimat[i:akhir])
	}

	return potongan
}

func perhitunganPotongan(potongan string, mu *sync.Mutex, frekuensi map[rune]int) {
	for _, huruf := range potongan {
		if isHuruf(huruf) {
			mu.Lock()
			frekuensi[huruf]++
			mu.Unlock()
		}
	}
}

func isHuruf(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func tampilkanFrekuensi(frekuensi map[rune]int) {
	fmt.Println("Frekuensi Huruf:")
	for huruf, jumlah := range frekuensi {
		fmt.Printf("%c: %d\n", huruf, jumlah)
	}
}