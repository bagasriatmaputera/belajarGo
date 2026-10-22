package main

import (
	"fmt"
)

func main() {
	var name string
	name = "Testing"
	age := 30
	// age := 40 -> ini jelas error karena sudah ada variable define age sehingga tidak perlu :=
	age = 40
	fmt.Println(name)
	fmt.Println(age)
}
