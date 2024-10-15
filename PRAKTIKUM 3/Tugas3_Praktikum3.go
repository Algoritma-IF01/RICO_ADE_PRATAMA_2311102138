// TUGAS 3 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM BIAYA POS UNTUK MENGHITUNG BIAYA PENGIRIMAN
package main

import (
	"fmt"
)

func main() {
	for i_2138 := 1; i_2138 <= 3; i_2138++ {
		var BeratParsel_2318, kg_2138, gram_2138 int
		var BiayaPerkg_2138, BiayaSisa_2138, TotalBiaya_2138 int

		// Input berat parsel dalam gram
		fmt.Print("Berat parsel (gram): ")
		fmt.Scan(&BeratParsel_2318)

		// Menghitung berat dalam kg dan sisa gram
		kg_2138 = BeratParsel_2318 / 1000   // Berat dalam kilogram
		gram_2138 = BeratParsel_2318 % 1000 // Sisa berat dalam gram

		// Menghitung biaya per kg
		BiayaPerkg_2138 = kg_2138 * 10000 // Biaya per kg adalah Rp. 10.000 per kg

		// Menghitung biaya sisa berdasarkan gram
		if kg_2138 >= 10 {
			// Jika berat lebih dari atau sama dengan 10 kg, biaya sisa gram diabaikan (gratis)
			BiayaSisa_2138 = 0
			gram_2138 = 0 // Mengatur gram menjadi 0 untuk mencerminkan gratisnya biaya
		} else {
			// Jika berat kurang dari 10 kg, hitung biaya berdasarkan sisa gram
			if gram_2138 <= 500 {
				BiayaSisa_2138 = gram_2138 * 5 // Jika gram <= 500, biaya Rp. 5 per gram
			} else {
				BiayaSisa_2138 = gram_2138 * 15 // Jika gram > 500, biaya Rp. 15 per gram
			}
		}

		// Menghitung total biaya
		TotalBiaya_2138 = BiayaPerkg_2138 + BiayaSisa_2138 // Total biaya adalah biaya per kg + biaya sisa gram

		// Output hasil perhitungan
		fmt.Printf("Detail berat: %d kg + %d gram\n", kg_2138, gram_2138)              // Menampilkan berat dalam kg dan gram
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", BiayaPerkg_2138, BiayaSisa_2138) // Menampilkan rincian biaya
		fmt.Printf("Total biaya: Rp. %d\n\n", TotalBiaya_2138)                         // Menampilkan total biaya
	}
}
