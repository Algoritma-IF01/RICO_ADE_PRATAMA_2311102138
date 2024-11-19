// TUGAS 3 : PRAKTIKUM 6
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM POS PELAYANAN TERPADU(POSYANDU) UNTUK MENCATAT DATA BERAT BALITA(DALAM KG)
package main

import (
	"fmt" // Mengimpor paket fmt untuk I/O
)

// Deklarasi tipe array untuk menyimpan berat balita
type arrBalita [100]float64 // Array untuk berat balita dengan kapasitas maksimum 100

// Fungsi untuk menghitung berat minimum dan maksimum dalam array
func hitungMinMax(arrBerat_2138 arrBalita, n_2138 int, bMin_2138, bMax_2138 *float64) {
	*bMin_2138 = arrBerat_2138[0] // Inisialisasi nilai minimum dengan elemen pertama array
	*bMax_2138 = arrBerat_2138[0] // Inisialisasi nilai maksimum dengan elemen pertama array

	// Perulangan untuk membandingkan nilai array dengan minimum dan maksimum saat ini
	for i_2138 := 1; i_2138 < n_2138; i_2138++ { // Mulai dari indeks kedua karena elemen pertama sudah diproses
		if arrBerat_2138[i_2138] < *bMin_2138 { // Jika berat balita lebih kecil dari minimum saat ini
			*bMin_2138 = arrBerat_2138[i_2138] // Perbarui nilai minimum
		}
		if arrBerat_2138[i_2138] > *bMax_2138 { // Jika berat balita lebih besar dari maksimum saat ini
			*bMax_2138 = arrBerat_2138[i_2138] // Perbarui nilai maksimum
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRata(arrBerat_2138 arrBalita, n_2138 int) float64 {
	total_2138 := 0.0 // Variabel untuk menyimpan total_2138 berat semua balita

	for i_2138 := 0; i_2138 < n_2138; i_2138++ { // Perulangan untuk menjumlahkan semua elemen array
		total_2138 += arrBerat_2138[i_2138] // Tambahkan berat balita ke total_2138
	}

	return total_2138 / float64(n_2138) // Hitung rata-rata dengan membagi total_2138 dengan jumlah balita
}

func main() {
	var n_2138 int                   // Variabel untuk jumlah balita
	var beratBalita_2138 arrBalita   // Array untuk menyimpan berat balita
	var bMin_2138, bMax_2138 float64 // Variabel untuk menyimpan nilai minimum dan maksimum

	// Input jumlah data berat balita
	fmt.Print("Masukkan banyak data berat balita: ") // Prompt pengguna untuk memasukkan jumlah balita
	fmt.Scan(&n_2138)                                // Membaca jumlah data balita dari input pengguna

	// Validasi input jumlah balita agar tidak melebihi kapasitas maksimum
	if n_2138 <= 0 || n_2138 > 100 { // Cek apakah jumlah balita valid (1 <= n_2138 <= 100)
		fmt.Println("Jumlah balita harus antara 1 dan 100.") // Pesan error jika input tidak valid
		return                                               // Keluar dari program jika input tidak valid
	}

	// Input berat balita
	fmt.Println("Masukkan berat balita:")        // Prompt untuk memasukkan berat setiap balita
	for i_2138 := 0; i_2138 < n_2138; i_2138++ { // Perulangan untuk membaca berat masing-masing balita
		fmt.Printf("Berat balita ke-%d: ", i_2138+1) // Tampilkan nomor balita
		fmt.Scan(&beratBalita_2138[i_2138])          // Baca berat balita dari input pengguna
	}

	// Panggil fungsi untuk menghitung nilai minimum dan maksimum
	hitungMinMax(beratBalita_2138, n_2138, &bMin_2138, &bMax_2138) // Menghitung nilai minimum dan maksimum

	// Panggil fungsi untuk menghitung rata-rata berat balita
	rataRata := hitungRata(beratBalita_2138, n_2138) // Menghitung rata-rata berat balita

	// Output hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin_2138) // Tampilkan nilai minimum
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax_2138)  // Tampilkan nilai maksimum
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)  // Tampilkan nilai rata-rata
}
