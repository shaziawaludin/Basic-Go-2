package main

import (
	"fmt"
)

func main() {
	var nilai int
	fmt.Println("Masukkan Angka : ")
	fmt.Scanln(&nilai)
	if nilai%2 == 0 {
		fmt.Printf("%d Adalah Genap", nilai)
	} else {
		fmt.Printf("%d Adalah Ganjil", nilai)
	}
}
