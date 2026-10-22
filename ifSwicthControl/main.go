package main

import (
	"fmt"
)

func main() {
	age := 11
	adult := 18
	fmt.Println("Ini If Case")
	if age >= adult {
		fmt.Println("Sudah Dewasa")
	} else {
		fmt.Println("Belum Dewasa")
	}

	fmt.Println("Ini Switch Case")
	score := 90
	switch score {
	case 100:
		fmt.Println("Nilai A")
	case 90:
		fmt.Println("Nilai B")
		fallthrough //ini akan tetap maksa ke case berikutnya
	case 80:
		fmt.Println("Nilai C")
	default:
		fmt.Println("Nilai D")
	}
}
