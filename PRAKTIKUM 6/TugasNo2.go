// TUGAS 2 : PRAKTIKUM 6
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM UNTUK MENENTUKAN TARIF IKAN YANG AKAN DIJUAL KE PASAR
package main

import (
	"fmt" // Mengimpor paket fmt untuk I/O
)

const maxKapasitas_2138 = 1000 // Konstanta untuk kapasitas maksimum jumlah ikan

// Fungsi untuk menghitung total berat ikan di setiap wadah dan rata-rata berat per wadah
func hitungBerat_2138(x_2138 int, y_2138 int, beratIkan_2138 []float64) ([]float64, float64) {
	// Hitung jumlah wadah yang diperlukan
	jumlahWadah_2138 := (x_2138 + y_2138 - 1) / y_2138   // Pembulatan ke atas jika ada sisa ikan
	beratWadah_2138 := make([]float64, jumlahWadah_2138) // Slice untuk menyimpan berat total setiap wadah

	// Distribusi berat ikan ke setiap wadah
	for i_2138 := 0; i_2138 < x_2138; i_2138++ { // Iterasi untuk semua ikan
		indeksWadah_2138 := i_2138 / y_2138                         // Tentukan indeks wadah berdasarkan kapasitas
		beratWadah_2138[indeksWadah_2138] += beratIkan_2138[i_2138] // Tambahkan berat ikan ke wadah yang sesuai
	}

	// Hitung rata-rata berat ikan per wadah
	totalBerat_2138 := 0.0 // Variabel untuk menyimpan total berat semua ikan
	for _, berat_2138 := range beratWadah_2138 {
		totalBerat_2138 += berat_2138 // Tambahkan berat wadah ke total
	}
	rataRataBerat_2138 := totalBerat_2138 / float64(jumlahWadah_2138) // Hitung rata-rata berat ikan

	return beratWadah_2138, rataRataBerat_2138 // Kembalikan hasil berupa total berat tiap wadah dan rata-rata berat
}

func main() {
	var x_2138, y_2138 int // Variabel untuk jumlah ikan dan kapasitas wadah

	// Input jumlah ikan (x) dan kapasitas wadah (y)
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x_2138, &y_2138)

	// Validasi input
	if x_2138 <= 0 || y_2138 <= 0 { // Pastikan jumlah ikan dan kapasitas wadah valid
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0 dan tidak boleh negatif.")
		return // Keluar dari program jika input tidak valid
	}

	if x_2138 > maxKapasitas_2138 { // Validasi maksimum jumlah ikan
		fmt.Printf("Jumlah ikan tidak boleh lebih dari %d.\n", maxKapasitas_2138)
		return // Keluar dari program jika jumlah ikan melebihi kapasitas maksimum
	}

	// Input berat masing-masing ikan
	beratIkan_2138 := make([]float64, x_2138) // Slice untuk menyimpan berat tiap ikan
	fmt.Printf("Masukkan berat ikan sebanyak %d:\n", x_2138)
	for i_2138 := 0; i_2138 < x_2138; i_2138++ { // Iterasi untuk input berat ikan
		fmt.Printf("Berat ikan ke-%d: ", i_2138+1)
		fmt.Scan(&beratIkan_2138[i_2138]) // Baca berat ikan
		if beratIkan_2138[i_2138] < 0 {   // Validasi berat ikan tidak boleh negatif
			fmt.Println("Berat ikan tidak boleh negatif.")
			return // Keluar dari program jika ada berat negatif
		}
	}

	// Panggil fungsi untuk menghitung total berat dan rata-rata
	beratWadah_2138, rataRataBerat_2138 := hitungBerat_2138(x_2138, y_2138, beratIkan_2138)

	// Output total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah adalah:")
	for i_2138, berat_2138 := range beratWadah_2138 {
		fmt.Printf("Wadah %d: %.2f kg\n", i_2138+1, berat_2138) // Tampilkan berat tiap wadah
	}

	// Output rata-rata berat ikan per wadah
	fmt.Printf("\nBerat rata-rata ikan di setiap wadah adalah: %.2f kg\n", rataRataBerat_2138)
}
