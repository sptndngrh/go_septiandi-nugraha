// Soal 4
// Kamu memiliki tiga bilangan bulat yang berbeda, x, y dan z, yang memenuhi tiga hubungan berikut:
// a. x + y + z = A
// b. xyz = B
// c. x^2 + y^2 + z^2 = C
// Kamu diminta untuk menulis sebuah program yang memecahkan x, y dan z untuk nilai yang diberikan A, B dan C. (1 ≤ A, B, C ≤ 10000).
// Sample test cases
// Input : 1 2 3
// Output : No Solution
// Input : 6 6 14
// Output : 1 2 3

package main

import "fmt"


func SimpleEquations(a, b, c int) {
	var x, y, z int

	for x = 1; x <= 10000; x++ {
		for y = 1; y <= 10000; y++ {
			z = a - x - y
			if z > 0 && x*y*z == b && x*x+y*y+z*z == c {
				break
			}
		}
		if y <= 10000 {
			break
		}
	}

	if x <= 10000 && y <= 10000 {
		fmt.Printf("%d %d %d\n", x, y, z)
	} else {
		fmt.Println("No solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3) // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}