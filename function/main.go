package main

import "fmt"

// function biasa
func sapa() {
	fmt.Println("Halo, nama saya ")
}

// function dengan parameter
func sapaNama(nama string) {
	fmt.Println("Halo, nama saya ", nama)
}

// function dengan parameter dan return value
func sapaNamaKembalikan(nama string) string {
	return "Halo, nama saya " + nama
}

// function dengan multiple return value
func sapaNamaUmur(nama string, umur int) (string, int) {
	return "Halo, nama saya " + nama, umur
}

// function dengan variadic parameter
func sapaVariadic(nama string, umur ...int) string {
	return "Halo, nama saya " + nama
}

func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa membagi dengan 0")
	}
	return a / b, nil
}

func main() {
	// Bagi
	hasil, err := bagi(10, 0)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Hasil : ", hasil)
	}

}
