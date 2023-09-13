// Soal 2
// Buat sebuah struct dengan nama Student yang mempunyai properti name dan score dalam bentuk
// slice kemudian simpan data siswa sebanyak 5 siswa. Setelah 5 siswa dimasukkan maka program menunjukkan
// skor rata-rata, siswa yang memiliki skor minimum dan maksimal! (Implementasikan penggunaan method).

package main

import (
	"fmt"
	"math"
)

type student struct {
	name string
	score float64
}

type listStudent []student

func (s listStudent) avgScore() float64 {
	totalScore := 0.0
	for _, student := range s {
		totalScore += student.score
	}
	return totalScore / float64(len(s))
}

func (s listStudent) maxScore() student {
	studentMaxscore := -1.0
	var maxStd student

	for _, student := range s {
		if student.score > studentMaxscore {
			studentMaxscore = student.score
			maxStd = student
		}
	}
	return maxStd
}

func (s listStudent) minScore() student {
	studentMinscore := math.MaxFloat64
	var minStd student

	for _, student := range s {
		if student.score < studentMinscore {
			studentMinscore = student.score
			minStd = student
		}
	}
	return minStd
}

func main() {
	var students listStudent

	for i := 1; i <= 5; i++ {
		var name string
		var score float64

		fmt.Printf("input %d student's Name : ", i)
		fmt.Scanln(&name)
		fmt.Printf("input %d student's Score : ", i)
		fmt.Scanln(&score)

		st := student{name, score}
		students = append(students, st)
	}

	avgScore := students.avgScore()
	minScore := students.minScore()
	maxScore := students.maxScore()

	fmt.Printf("average score: %.2f\n", avgScore)
	fmt.Printf("min score of student; %s (%.0f)\n", minScore.name, minScore.score)
	fmt.Printf("max score of student: %s (%.0f)\n", maxScore.name, maxScore.score)
}