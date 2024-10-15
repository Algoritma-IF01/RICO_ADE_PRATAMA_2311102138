// TUGAS 5 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MENCARI DAN MENAMPILKAN SEMUA FAKTOR BILANGAN, KEMUDIAN MENENTUKAN APAKAH B MERUPAKAN BILANGAN PRIMA
package main

import (
	"fmt"
)

func main() {
	var b_2138 int // Variabel untuk menyimpan bilangan yang dimasukkan pengguna

	// Input bilangan bulat dari pengguna
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&b_2138)

	// Mencari dan menampilkan faktor-faktor dari bilangan b_2138
	fmt.Print("Faktor: ")
	HitungFaktor_2138 := 0 // Variabel untuk menghitung jumlah faktor (digunakan untuk menentukan bilangan prima)
	for i_2138 := 1; i_2138 <= b_2138; i_2138++ {
		if b_2138%i_2138 == 0 {
			fmt.Print(i_2138, " ") // Menampilkan faktor
			HitungFaktor_2138++    // Menambah jumlah faktor
		}
	}
	fmt.Println()

	// Menentukan apakah bilangan b_2138 adalah bilangan prima
	if HitungFaktor_2138 == 2 {
		fmt.Println("Prima: true") // Bilangan prima hanya memiliki 2 faktor: 1 dan dirinya sendiri
	} else {
		fmt.Println("Prima: false") // Jika faktor lebih dari 2, berarti bukan bilangan prima
	}
}
