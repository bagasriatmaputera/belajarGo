package main

import (
	"fmt"
)

func main() {

	// cara deklarasi slice yang pertama
	var slice []string
	slice = append(slice, "Bagas")
	slice = append(slice, "Putra")
	slice = append(slice, "Budi")

	// cara pendeklarasian slice yang lebih ringkas
	slice2 := []string{
		"A",
		"B",
		"C",
	}
	fmt.Println(slice)
	fmt.Println(slice2)

	for _, data := range slice {
		fmt.Println(data)
	}

	for _, data := range slice2 {
		fmt.Println(data)
	}
}
