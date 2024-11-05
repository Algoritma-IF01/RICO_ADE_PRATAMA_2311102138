// BUKBER IF TYPE K
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM PENENTUAN MENU TAKJIL
package main

import (
	"fmt"
)

func main() {
	// Masukkan jumlah peserta
	var n_2138 int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n_2138)

	// Deklarasi variabel untuk menyimpan jumlah yang mendapatkan es tebak, es cendol, dan lamang
	countTebak_2138, countCendol_2138, countLamang_2138 := 0, 0, 0

	// Iterasi untuk setiap peserta
	for i_2138 := 0; i_2138 < n_2138; i_2138++ {
		var cardNumber_2138 int
		fmt.Printf("Masukkan nomor kartu peserta %d: ", i_2138+1)
		fmt.Scan(&cardNumber_2138)

		// Cek digit dari nomor kartu
		allOdd_2138, allEven_2138 := true, true
		for cardNumber_2138 > 0 {
			digit_2138 := cardNumber_2138 % 10
			if digit_2138%2 == 0 {
				allOdd_2138 = false
			} else {
				allEven_2138 = false
			}
			cardNumber_2138 /= 10
		}

		// Tentukan menu tajil berdasarkan digit nomor kartu
		if allOdd_2138 {
			fmt.Println("Es Tebak")
			countTebak_2138++
		} else if allEven_2138 {
			fmt.Println("Es Cendol")
			countCendol_2138++
		} else {
			fmt.Println("Lamang")
			countLamang_2138++
		}
	}

	// Tampilkan hasil jumlah peserta yang mendapatkan setiap tajil
	fmt.Printf("Jumlah peserta yang mendapatkan Es Tebak: %d\n", countTebak_2138)
	fmt.Printf("Jumlah peserta yang mendapatkan Es Cendol: %d\n", countCendol_2138)
	fmt.Printf("Jumlah peserta yang mendapatkan Lamang: %d\n", countLamang_2138)
}
