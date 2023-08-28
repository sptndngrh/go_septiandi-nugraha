// Soal 1
// Buatlah sebuah method untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi 
// (1.5 L / mill), method tersebut dimiliki oleh struct Car yang memiliki property carType dan fuelin.
// a. Input: fuelin: 10.5, carType: "SUV"
// b. Output: car type:  SUV , est. distance:  7

package main

import "fmt"

type carType struct {
	carType string
	fuelin float64
}

func (e carType) calculateDistance() float64 {
	return e.fuelin * 1.5
}

func main() {
	car := carType{"SUV", 10.5}
	perkiraanBahanbakar := car.calculateDistance()
	fmt.Printf("car type: %s , est. distance: %.2f\n", car.carType, perkiraanBahanbakar)
}