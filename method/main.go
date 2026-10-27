package main

import (
	"fmt"
)

type Persegi struct {
	Sisi int
}

func main() {

	p := Persegi{Sisi: 10}
	fmt.Println("Sebelum di ubah :", p.Sisi)
	p.ubahNonPointer()
	fmt.Println("Sesudah di ubah (method value) :", p.Sisi)
	p.ubahPointer()
	fmt.Println("Sesudah di ubah (method pointer) :", p.Sisi)
	fmt.Println("Informasi Persegi= ", p.Info())
}

// method value
func (p Persegi) Luas() int {
	return p.Sisi * p.Sisi
}

func (p Persegi) Keliling() int {
	return 4 * p.Sisi
}

func (p Persegi) Info() string {
	return fmt.Sprintf("luas=%d keliling=%d", p.Luas(), p.Keliling())
}

// method pointer
func (p *Persegi) ubahPointer() {
	p.Sisi = 30
}

// method value
func (p Persegi) ubahNonPointer() {
	p.Sisi = 20
}
