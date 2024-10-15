// TUGAS 4 C : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MENERIMA INPUT SEBUAH BILANGAN REAL YANG MENYATAKAN NAM, MENGHITUNG NMK DAN MENAMPILKANNYA
package main

import (
	"fmt"
)

func main() {
	var nam_2138 float64 // Variabel untuk menyimpan nilai akhir mata kuliah (NAM)
	var nmk_2138 string  // Variabel untuk menyimpan nilai mutu kuliah (NMK)

	// Input nilai akhir mata kuliah dari pengguna
	fmt.Print("Nilai akhir mata kuliah (NAM): ")
	fmt.Scanln(&nam_2138) // Membaca input nilai NAM

	// Menentukan NMK berdasarkan nilai NAM menggunakan percabangan
	if nam_2138 >= 80 {
		// Jika nilai lebih besar atau sama dengan 80, NMK adalah A
		nmk_2138 = "A"
	} else if nam_2138 >= 72.5 {
		// Jika nilai antara 72.5 hingga 79.9, NMK adalah AB
		nmk_2138 = "AB"
	} else if nam_2138 >= 65 {
		// Jika nilai antara 65 hingga 72.4, NMK adalah B
		nmk_2138 = "B"
	} else if nam_2138 >= 57.5 {
		// Jika nilai antara 57.5 hingga 64.9, NMK adalah BC
		nmk_2138 = "BC"
	} else if nam_2138 >= 50 {
		// Jika nilai antara 50 hingga 57.4, NMK adalah C
		nmk_2138 = "C"
	} else if nam_2138 >= 40 {
		// Jika nilai antara 40 hingga 49.9, NMK adalah D
		nmk_2138 = "D"
	} else {
		// Jika nilai di bawah 40, NMK adalah E
		nmk_2138 = "E"
	}

	// Output NMK yang telah ditentukan berdasarkan nilai NAM
	fmt.Println("Nilai mutu kuliah (NMK):", nmk_2138) // Menampilkan hasil NMK
}
