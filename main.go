package main

import (
	"fmt"
)

func main() {
	var namaKakak string = "merosa pesona"
	fmt.Println("Nama Kakak:", namaKakak)
	fmt.Print("Nama Adik: ")
	kakak := len(namaKakak)
	for i := kakak - 1; i >= 0; i-- {
		fmt.Print(string(namaKakak[i]))
	}

}
