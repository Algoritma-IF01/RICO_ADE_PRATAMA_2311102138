// TUGAS 3 : PRAKTIKUM 5
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM UNTUK MEREKAP SKOR PERTANDINGAN 2 BUAH KLUB BOLA YANG BERLAGA
package main

import "fmt" // Mengimpor paket fmt untuk I/O

func main() {
	// Deklarasi variabel untuk nama klub A dan klub B
	var klubA_2138, klubB_2138 string
	// Deklarasi variabel untuk skor masing-masing klub
	var scoreA_2138, scoreB_2138 int
	// Slice untuk menyimpan hasil pertandingan
	hasil_2138 := []string{}
	// Variabel untuk menghitung jumlah pertandingan
	jumlahMatch_2138 := 1

	// Input nama klub A
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA_2138)
	// Input nama klub B
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB_2138)

	// Loop untuk input skor pertandingan
	for {
		// Input skor pertandingan untuk masing-masing klub
		fmt.Printf("Pertandingan %d: ", jumlahMatch_2138)
		fmt.Scan(&scoreA_2138, &scoreB_2138)

		// Jika skor negatif, keluar dari loop
		if scoreA_2138 < 0 || scoreB_2138 < 0 {
			break
		}

		// Menentukan pemenang atau draw
		if scoreA_2138 > scoreB_2138 {
			hasil_2138 = append(hasil_2138, klubA_2138)
		} else if scoreB_2138 > scoreA_2138 {
			hasil_2138 = append(hasil_2138, klubB_2138)
		} else {
			hasil_2138 = append(hasil_2138, "Draw")
		}

		// Tambah jumlah pertandingan
		jumlahMatch_2138++
	}

	// Menampilkan hasil pertandingan
	fmt.Println("\nDaftar hasil pertandingan:")
	for i_2138, poin_2138 := range hasil_2138 {
		if poin_2138 == "Draw" {
			fmt.Printf("Hasil %d: Draw\n", i_2138+1)
		} else {
			fmt.Printf("Hasil %d: %s\n", i_2138+1, poin_2138)
		}
	}

	// Menampilkan pesan bahwa pertandingan selesai
	fmt.Println("Pertandingan selesai")
}
