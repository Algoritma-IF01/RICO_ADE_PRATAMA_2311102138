// TUGAS 2 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MENERIMA INPUT BILANGAN SEBUAH K
package main

import (
	"fmt"
	"math"
)

// Menghitung nilai fungsi f(k_2138) berdasarkan rumus yang diberikan
func f(k_2138 float64) float64 {
	// Menghitung pembilang: (4*k_2138 + 2)^2
	Numerator_2138 := math.Pow((4*k_2138 + 2), 2)
	// Menghitung penyebut: (4*k_2138 + 1)*(4*k_2138 + 3)
	Denominator_2138 := (4*k_2138 + 1) * (4*k_2138 + 3)
	// Mengembalikan nilai f(k_2138)
	return Numerator_2138 / Denominator_2138
}

// Menghitung nilai √2 dengan menjumlahkan f(k_2138) dari k_2138=0 hingga k_2138 yang diberikan
func sqrt2(k_2138 int) float64 {
	// Inisialisasi hasil dengan 1
	Hasil_2138 := 1.0
	// Loop untuk menghitung hasil
	for i := 0; i <= k_2138; i++ {
		// Mengalikan hasil dengan f(i)
		Hasil_2138 *= f(float64(i))
	}
	// Mengembalikan nilai hasil akhir
	return Hasil_2138
}

func main() {
	// Deklarasi variabel untuk menyimpan nilai K
	var K_2138 int

	// Loop untuk meminta input 3 kali dari pengguna
	for i := 1; i <= 3; i++ {
		// Meminta pengguna memasukkan nilai K
		fmt.Print("Nilai K = ")
		fmt.Scan(&K_2138)

		// Menghitung nilai akar 2 berdasarkan K yang diberikan
		ApproxSqrt2_1238 := sqrt2(K_2138)
		// Menampilkan hasil dengan format 10 angka di belakang koma
		fmt.Printf("Nilai akar 2 = %.10f\n\n", ApproxSqrt2_1238)
	}
}
