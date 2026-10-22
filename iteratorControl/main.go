package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println("Ini For Control")
		fmt.Println(i)
	}

	i := 1
	for i <= 20 {
		fmt.Println("Ini For ", i)
		i++
	}

	arrayString := []string{"Bagas", "Putra", "Budi"}
	fmt.Println("Ini For Range Hanya Nilai")
	for _, data := range arrayString { // _ itu mengabaikkan index
		fmt.Println("Data :", data)
	}

	fmt.Println("Ini For Range Nilai dan Index")
	name := "Bagas"
	for idx, data := range name {
		fmt.Println("Index :", idx, "Data :", string(data)) // data itu bentuknya byte jadi perlu diubah ke string -> string()
	}

	umur := map[string]int{"Budi": 20, "Sari": 22} //tipe data: map dengan key bertipe string dan value bertipe int

	for nama, u := range umur {
		fmt.Println(nama, u)
	}
}
