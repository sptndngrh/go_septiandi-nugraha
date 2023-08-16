package ganjilgenap

import "fmt"

func ganjilgenap() {
	var bil int

	fmt.Print("Masukkan bilangan yang anda kehendaki: ")
	fmt.Scanln(&bil)

	if bil%2 == 0 {
		fmt.Printf("%d merupakan bilangan genap. \n", bil)
	} else {
		fmt.Printf("%d merupakan bilangan ganjil. \n", bil)
	}
}