// TUGAS 1 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM PAK ANDI MENERIMA INPUT DUA BILANGAN REAL POSITIF
package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi variabel untuk menyimpan berat kantong
	var KantongKiri_2138, KantongKanan_2138 float64

	// Loop untuk meminta input berat dari pengguna
	for {
		// Meminta pengguna memasukkan berat belanjaan di kedua kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&KantongKiri_2138, &KantongKanan_2138)

		// Cek apakah salah satu kantong negatif atau total berat melebihi 150 kg
		if KantongKiri_2138 < 0 || KantongKanan_2138 < 0 || KantongKiri_2138+KantongKanan_2138 > 150 {
			// Jika ya, tampilkan pesan dan keluar dari loop
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat antara kantong kiri dan kanan
		Selisih_2138 := math.Abs(KantongKiri_2138 - KantongKanan_2138)

		// Tentukan apakah sepeda motor Pak Andi akan oleng
		Oleng_2138 := Selisih_2138 >= 9

		// Tampilkan hasil apakah sepeda motor oleng
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", Oleng_2138)
	}
}
