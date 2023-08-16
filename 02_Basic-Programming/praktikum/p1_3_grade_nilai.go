package gradenilai

import (
	"fmt"
)

func gradenilai() {
	var nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)

	var peringkat string

	switch {
	case nilai >= 80 && nilai <= 100:
		peringkat = "A"
	case nilai >= 65 && nilai <= 79:
		peringkat = "B"
	case nilai >= 50 && nilai <= 64:
		peringkat = "C"
	case nilai >= 35 && nilai <= 49:
		peringkat = "D"
	case nilai >= 0 && nilai <= 34:
		peringkat = "E"
	default:
		peringkat = "Nilai Invalid"
	}

	fmt.Printf("Grade: %s\n", peringkat)
}