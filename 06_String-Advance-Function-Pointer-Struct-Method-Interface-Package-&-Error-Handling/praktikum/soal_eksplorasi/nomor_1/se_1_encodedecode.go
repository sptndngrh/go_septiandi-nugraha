// Soal 1
// Implementasikan interface yang terdiri dari method encode dan decode. Algoritma enkripsi yang dapat digunakan adalah subtitusi cipher.

package main

import "fmt"


type student struct {
	name string
	nameEncode string
	score int
}


type Chiper interface {
	Encode() string
	Decode() string
}


func (s *student) Encode() string {
	var nameEncode = ""
	
	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			encodedChar := 'a' + (char-'a'-9)%26
			nameEncode += string(encodedChar)
		} else if char >= 'A' && char <= 'Z' {
			encodedChar := 'A' + (char-'A'-9)%26
			nameEncode += string(encodedChar)
		} else {
			nameEncode += string(char)
		}
	}

	return nameEncode
}


func (s *student) Decode() string {
	var nameDecode = ""
	
	for _, char := range s.nameEncode {
		if char >= 'a' && char <= 'z' {
			decodedChar := 'a' + (char-'a'-9+26)%26
			nameDecode += string(decodedChar)
		} else if char >= 'A' && char <= 'Z' {
			decodedChar := 'A' + (char-'a'-9+26)%26
			nameDecode += string(decodedChar)
		} else {
			nameDecode += string(char)
		}
	}

	return nameDecode
}


func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student's name " + a.name + " is : " + c.Decode())
	}

}