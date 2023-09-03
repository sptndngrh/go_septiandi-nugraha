// Soal 1
// Buatlah sebuah fungsi yang mencetak angka kelipatan x, dimana x merupakan parameter bilangan bulat positif.
// lalu jalankan fungsi tersebut dengan menerapkan goroutine, dengan interval waktu 3 detik!

package main

import (
	"fmt"
	"time"
)

func printMultiple(x int) {
	for i := 1; ; i++ {
		fmt.Printf("%d kelipatan adalah %d\n", i, x*i)
		const durationMultipleSecond = 3
		time.Sleep(durationMultipleSecond * time.Second)
	}
}

func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat (positif): ")
	_, err := fmt.Scan(&x)
	if err != nil || x <= 0 {
		fmt.Println("Masukkan bilangan bulat positif!")
		return
	}
	go printMultiple(x)

	const sleepDurationSeconds = 10
	time.Sleep(sleepDurationSeconds * time.Second)
}