// TUGAS 4 : PRAKTIKUM 7
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM PERPUSTAKAAN UNTUK MENGELOLA DATA BUKU DI DALAM SUATU PERPUSTAKAAN MENGGUNAKAN INSERTION SORT
package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

const nMax_2138 = 7919 // Menentukan jumlah maksimum buku yang dapat ditampung

// Struct Buku_2138 untuk menyimpan informasi tentang buku
type Buku_2138 struct {
	id_2138       int    // ID buku
	judul_2138    string // Judul buku
	penulis_2138  string // Penulis buku
	penerbit_2138 string // Penerbit buku
	tahun_2138    int    // Tahun terbit
	rating_2138   int    // Rating buku
}

// Fungsi untuk mengisi data buku ke dalam pustaka
func DaftarkanBuku_2138(pustaka_2138 []Buku_2138, n_2138 int) []Buku_2138 {
	reader_2138 := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input
	for i_2138 := 0; i_2138 < n_2138; i_2138++ {
		fmt.Printf("\nMasukkan data untuk buku ke-%d:\n", i_2138+1)

		// Membaca ID buku
		fmt.Print("ID: ")
		idInput_2138, _ := reader_2138.ReadString('\n')
		idInput_2138 = strings.TrimSpace(idInput_2138)
		id_2138, _ := strconv.Atoi(idInput_2138)

		// Membaca Judul buku
		fmt.Print("Judul: ")
		judul_2138, _ := reader_2138.ReadString('\n')
		judul_2138 = strings.TrimSpace(judul_2138)

		// Membaca Penulis buku
		fmt.Print("Penulis: ")
		penulis_2138, _ := reader_2138.ReadString('\n')
		penulis_2138 = strings.TrimSpace(penulis_2138)

		// Membaca Penerbit buku
		fmt.Print("Penerbit: ")
		penerbit_2138, _ := reader_2138.ReadString('\n')
		penerbit_2138 = strings.TrimSpace(penerbit_2138)

		// Membaca Tahun buku
		fmt.Print("Tahun: ")
		tahunInput_2138, _ := reader_2138.ReadString('\n')
		tahunInput_2138 = strings.TrimSpace(tahunInput_2138)
		tahun_2138, _ := strconv.Atoi(tahunInput_2138)

		// Membaca Rating buku
		fmt.Print("Rating: ")
		ratingInput_2138, _ := reader_2138.ReadString('\n')
		ratingInput_2138 = strings.TrimSpace(ratingInput_2138)
		rating_2138, _ := strconv.Atoi(ratingInput_2138)

		// Menambahkan data ke dalam pustaka
		pustaka_2138[i_2138] = Buku_2138{id_2138, judul_2138, penulis_2138, penerbit_2138, tahun_2138, rating_2138}
	}
	return pustaka_2138
}

// Fungsi untuk mengurutkan buku berdasarkan rating secara menurun menggunakan Insertion Sort
func UrutkanBuku_2138(pustaka_2138 []Buku_2138, n_2138 int) {
	for i_2138 := 1; i_2138 < n_2138; i_2138++ {
		key_2138 := pustaka_2138[i_2138] // Elemen yang akan ditempatkan
		j_2138 := i_2138 - 1

		// Geser elemen yang rating-nya lebih kecil
		for j_2138 >= 0 && pustaka_2138[j_2138].rating_2138 < key_2138.rating_2138 {
			pustaka_2138[j_2138+1] = pustaka_2138[j_2138]
			j_2138--
		}
		pustaka_2138[j_2138+1] = key_2138 // Letakkan elemen di tempat yang sesuai
	}
}

// Fungsi untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit_2138(pustaka_2138 []Buku_2138, n_2138 int) {
	if n_2138 == 0 { // Jika tidak ada buku
		fmt.Println("Tidak ada buku di perpustakaan.") // Tampilkan pesan
		return
	}

	idxTerfavorit_2138 := 0                             // Indeks buku dengan rating tertinggi
	ratingTertinggi_2138 := pustaka_2138[0].rating_2138 // Ambil rating dari buku pertama

	// Iterasi untuk mencari buku dengan rating tertinggi
	for i_2138 := 1; i_2138 < n_2138; i_2138++ {
		if pustaka_2138[i_2138].rating_2138 > ratingTertinggi_2138 {
			ratingTertinggi_2138 = pustaka_2138[i_2138].rating_2138
			idxTerfavorit_2138 = i_2138
		}
	}

	// Tampilkan buku terfavorit dengan rating tertinggi
	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		pustaka_2138[idxTerfavorit_2138].judul_2138, pustaka_2138[idxTerfavorit_2138].penulis_2138,
		pustaka_2138[idxTerfavorit_2138].penerbit_2138, pustaka_2138[idxTerfavorit_2138].tahun_2138,
		pustaka_2138[idxTerfavorit_2138].rating_2138)
}

// Fungsi untuk mencetak lima buku dengan rating tertinggi
func CetakSetTerbaru_2138(pustaka_2138 []Buku_2138, n_2138 int) {
	fmt.Println("\nTOP 5 buku dengan rating tertinggi:")
	// Pastikan buku diurutkan sebelum mencetak
	UrutkanBuku_2138(pustaka_2138, n_2138)
	for i_2138 := 0; i_2138 < 5 && i_2138 < n_2138; i_2138++ {
		fmt.Printf("Judul: %s, Rating: %d\n", pustaka_2138[i_2138].judul_2138, pustaka_2138[i_2138].rating_2138)
	}
}

// Fungsi untuk mencari buku berdasarkan rating tertentu
func CariBuku_2138(pustaka_2138 []Buku_2138, n_2138 int, ratingCari_2138 int) {
	ditemukan_2138 := false // Penanda jika buku ditemukan

	// Cari buku dengan rating yang sesuai
	fmt.Printf("\nBuku dengan rating %d:\n", ratingCari_2138)
	for i_2138 := 0; i_2138 < n_2138; i_2138++ {
		if pustaka_2138[i_2138].rating_2138 == ratingCari_2138 {
			// Tampilkan data buku
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n\n",
				pustaka_2138[i_2138].judul_2138, pustaka_2138[i_2138].penulis_2138,
				pustaka_2138[i_2138].penerbit_2138, pustaka_2138[i_2138].tahun_2138,
				pustaka_2138[i_2138].rating_2138)
			ditemukan_2138 = true
		}
	}

	// Jika tidak ada buku yang ditemukan
	if !ditemukan_2138 {
		fmt.Printf("Tidak ada buku dengan rating %d\n", ratingCari_2138)
	}
}

// Program utama
func main() {
	reader_2138 := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input

	// Membaca jumlah buku di perpustakaan
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	nInput_2138, _ := reader_2138.ReadString('\n')
	nInput_2138 = strings.TrimSpace(nInput_2138)
	nPustaka_2138, _ := strconv.Atoi(nInput_2138)

	// Membuat slice untuk menampung data buku
	var Pustaka_2138 = make([]Buku_2138, nPustaka_2138)

	// Memanggil fungsi untuk mengisi data buku
	Pustaka_2138 = DaftarkanBuku_2138(Pustaka_2138, nPustaka_2138)
	fmt.Println()

	// Menampilkan buku dengan rating tertinggi
	CetakTerfavorit_2138(Pustaka_2138, nPustaka_2138)

	// Menampilkan lima buku dengan rating tertinggi
	CetakSetTerbaru_2138(Pustaka_2138, nPustaka_2138)

	// Mencari buku berdasarkan rating tertentu
	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	ratingCariInput_2138, _ := reader_2138.ReadString('\n')
	ratingCariInput_2138 = strings.TrimSpace(ratingCariInput_2138)
	ratingCari_2138, _ := strconv.Atoi(ratingCariInput_2138)
	CariBuku_2138(Pustaka_2138, nPustaka_2138, ratingCari_2138)
}
