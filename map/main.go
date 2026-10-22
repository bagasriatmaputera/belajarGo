package main

import (
	"fmt"
)

func main() {
	//  umur := map[string]int{"Budi": 20, "Sari": 22} // literal
	// m := make(map[string]int)                      // map kosong
	// var n map[string]int                           // nil map
	umur := map[string]int{"Budi": 20}

	umur["Sari"] = 22 // tambah
	umur["Budi"] = 21 // ubah

	// hapus data
	delete(umur, "Budi")
	delete(umur, "tidak-ada") // aman, tidak error

	fmt.Println(umur["Budi"]) // 21
	fmt.Println(len(umur))    // 2

	pemainBola := map[string]string{
		"nama":        "Budi",
		"posisi":      "Striker",
		"no_punggung": "10",
	}

	for idx, val := range pemainBola {
		fmt.Println("key: ", idx, "value: ", val)
	}

	sliceMap := []map[string]string{
		{
			"nama":   "Budi",
			"posisi": "Striker",
		},
		{
			"nama":   "Sari",
			"posisi": "Defender",
		},
	}

	for _, pemain := range sliceMap {
		fmt.Println("Nama : ", pemain["nama"])
		fmt.Println("Posisi : ", pemain["posisi"])
		fmt.Println("-------")
	}

}
