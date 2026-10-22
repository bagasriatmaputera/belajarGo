package main

import "fmt"

type user struct {
	Nama      string
	Umur      int
	Pekerjaan string
}

func main() {
	// cara 1
	userTest := user{"Bagas", 23, "Software Engineer"}

	// cara 2
	userTest1 := user{}
	userTest1.Nama = "Budi"
	userTest1.Umur = 22
	userTest1.Pekerjaan = "Web Developer"

	fmt.Println(userTest)
	fmt.Println(userTest1)

	data := []any{"Bagas Riatma Putera", 24, "Data Scientist"}
	updateUser(&userTest, data)

	fmt.Println("Setalah pakai func update: ", userTest)
}

// struct dalam function
// kenapa pakai tanda &? -> karena mau update value yang di struct asli
// jadi kita memakai pointer / alamat memori dari struct tersebut
func updateUser(user *user, data []any) {
	user.Nama = data[0].(string)
	user.Umur = int(data[1].(int))
	user.Pekerjaan = data[2].(string)
}

// struct dalam struct atau embeded struct
