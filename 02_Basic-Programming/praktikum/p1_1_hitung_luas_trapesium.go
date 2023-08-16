package hitungluas

import "fmt"

func hitungluas() {
	var p1, p2, t1 float64

	fmt.Print("      Program hitung luas trapesium      \n     masukkan satuan dalam satuan cm      \n")
	fmt.Print("                                         \n")
	fmt.Print("Masukkan sisi atas: ")
	fmt.Scanln(&p1)

	fmt.Print("Masukkan sisi bawah: ")
	fmt.Scanln(&p2)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&t1)

	//======================================
	fmt.Print("                                         \n")
	luasTrapesium := 0.5 * (p1 + p2) * t1
	fmt.Printf("Luas trapesium yaitu: %.2f cm", luasTrapesium)
}
