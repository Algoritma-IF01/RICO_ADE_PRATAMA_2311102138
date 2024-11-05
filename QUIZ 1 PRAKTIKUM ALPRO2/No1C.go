// SOAL 1 DIGIT TYPE J
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-D

// PROGRAM MEMOTONG SUATU BILANGAN BULAT POSITIF
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Menginput Bilangan Positif
	var input_2138 int
	fmt.Print("Masukkan bilangan bulat positif lebih besar dari 10: ")
	fmt.Scan(&input_2138)

	numStr_2138 := strconv.Itoa(input_2138)
	numLen_2138 := len(numStr_2138)

	var kiri_2138, kanan_2138 string

	// Potong menjadi 2 bagian
	if numLen_2138%2 == 0 {
		kiri_2138 = numStr_2138[:numLen_2138/2]
		kanan_2138 = numStr_2138[numLen_2138/2:]
	} else {
		kiri_2138 = numStr_2138[:numLen_2138/2+1]
		kanan_2138 = numStr_2138[numLen_2138/2+1:]
	}

	KiriNum_2138, _ := strconv.Atoi(kiri_2138)
	KananNum_2138, _ := strconv.Atoi(kanan_2138)

	sum_2138 := KiriNum_2138 + KananNum_2138
	fmt.Println(KiriNum_2138, KananNum_2138)
	fmt.Println(sum_2138)
}
