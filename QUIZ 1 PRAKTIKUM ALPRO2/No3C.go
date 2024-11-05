// PERKALIAN DENGAN CARA PENJUMLAHAN TYPE J
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-D

// MENGHITUNG JUMLAH PERTEMUAN 2 ORANG
package main

import "fmt"

// Fungsi rekursif untuk perkalian dengan penjumlahan
func multiply(n_2138, m_2138 int) int {
	// Basis rekursif
	if m_2138 == 0 {
		return 0
	} else if m_2138 == 1 {
		return n_2138
	}
	// Rekursi
	return n_2138 + multiply(n_2138, m_2138-1)
}

func main() {
	// Masukkan bilangan bulat n_2138 dan m
	var n_2138, m_2138 int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n_2138)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m_2138)

	// Panggil fungsi rekursif untuk menghitung hasil perkalian
	hasil_2138 := multiply(n_2138, m_2138)

	// Tampilkan hasil perkalian
	fmt.Printf("Hasil perkalian %d x %d adalah: %d\n", n_2138, m_2138, hasil_2138)
}
