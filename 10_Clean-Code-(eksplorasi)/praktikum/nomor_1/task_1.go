// Soal 1
// Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
// a. Berapa banyak kekurangan dalam penulisan kode tersebut?
// b. Bagian mana saja terjadi kekurangan tersebut?
// c. Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.

/*
package main

type user struct {
	id int
	username int
	password int
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}

a. Kekurangan dari code ini ada 6 poin yang diperhatikan

b. 	> Adanya package main tetapi package tersebut tidak fungsikan kembali dalam
   	fungsi main
	> Penamaan variabel untuk u, t dan r yang tidak jelas, tidak deskriptif dan sulit untuk
   	dipahami
	> Tipe data dengan nama 'user', 'userservice', 'getallusers' 'getuserbyid' harus diawali huruf besar seperti contoh
	'User', 'UserService', 'GetAllUser' 'GetUserByID' karena tipe data dengan format penamaan tersebut menandakan bahwa
	tipe data tersebut bisa diakses di luar package
	> Metode pointernya kurang konsisten, karena dalam code tersebut ada 'u uervice' yang dimana harusnya diubah
	menjadi 'User *UserService'
	> Komentar di code yang tidak lengkap untuk menjelaskan code tersebut itu memiliki fungsi apa dan dipahami oleh orang lain
	> Kurangnya error handling di code tersebut di bagian 'getuserbyid'

c. 	> Penamaan Variabel: Penamaan variabel yang tidak deskriptif membuat kode sulit dibaca dan dipahami. Variabel sebaiknya 
	  memiliki nama yang menjelaskan peran dan isinya.
	> Penamaan Metode: Konvensi penulisan metode dalam Go adalah menggunakan huruf besar pada awal kata (contoh: GetAllUsers, 
	  GetUserByID). Ini membantu dalam konsistensi dan pemahaman kode.
	> Penamaan Tipe Data: Tipe data yang tidak diawali huruf besar dapat membuat bingung antara variabel dan tipe data. Tipe 
	  data seharusnya memiliki penamaan yang lebih deskriptif dan sesuai dengan konvensi.
	> Kurangnya Error Handling: Tanpa error handling yang memadai, kode tidak dapat mengatasi situasi ketika pencarian pengguna 
	  gagal. Error handling membantu dalam memberikan umpan balik yang lebih baik kepada pengguna.
	> Komentar Tidak Lengkap: Komentar yang tidak lengkap membuat kode sulit dipahami oleh orang lain yang membacanya. 
	  Komentar sebaiknya memberikan penjelasan yang jelas tentang tujuan dan fungsionalitas kode.
	> Package Tidak Digunakan: Jika ada deklarasi package main, maka seharusnya ada fungsi main yang sesuai. Dalam kode 
	  asli, package main tidak memiliki fungsi main yang sesuai sehingga package tersebut tidak digunakan dengan benar.
*/

// Berikut code yang telah diperbaiki

package main

import (
	"errors"
	"fmt"
	"time"
)

// Pengguna dipresentasikan menjadi struktur User
type User struct {
	ID       int
	Username string
	Password string
}

// Operasi pengguna yang menyediakan layanan pengguna adalah UserService
type UserService struct {
	Users []User
}

// data pengguna dikembalikan dengan GetAllUser
func (userService *UserService) GetAllUser() []User {
	return userService.Users
}

// ID sebagai referensi untuk mengembalikan data pengguna dengan GetUserByID
func (userService *UserService) GetUserByID(id int) (User, error) {
	fmt.Println("Mencari data...")
	time.Sleep(2 * time.Second)

	for _, user := range userService.Users {
		if id == user.ID {
			return user, nil
		}
	}
	return User{}, errors.New("User not found")
}

func main() {
	userService := UserService{
		Users: []User{
			{ID: 1, Username: "Maharani", Password: "maharani123"},
			{ID: 2, Username: "Hakamoto", Password: "hakamoto123"},
			{ID: 3, Username: "Muhammad", Password: "mhmd123"},
		},
	}

	allUsers := userService.GetAllUser()
	fmt.Println("All Users: ")
	for _, user := range allUsers {
		fmt.Printf("ID: %d, Username: %s\n", user.ID, user.Username)
	}

	var UserID int
	fmt.Print("Masukkan ID pengguna yang ingin anda cari: ")
	fmt.Scan(&UserID)

	FoundUser, err := userService.GetUserByID(UserID)
	if err != nil {
		fmt.Printf("User with ID %d not found.\n", UserID)
	} else {
		fmt.Printf("User found - ID: %d, Username: %s\n", FoundUser.ID, FoundUser.Username)
	}
}
