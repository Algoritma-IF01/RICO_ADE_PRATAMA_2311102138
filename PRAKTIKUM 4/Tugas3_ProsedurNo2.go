// TUGAS 3 : PRAKTIKUM 4
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM GEMA MENCARI PEMENANG DARI DAFTAR PESERTA
package main

import (
	"bufio"   // Package untuk membaca input dari pengguna melalui terminal
	"fmt"     // Package untuk format I/O seperti print ke layar
	"os"      // Package untuk mengakses fungsi sistem operasi, seperti membaca input
	"strconv" // Package untuk melakukan konversi tipe data, seperti dari string ke integer
	"strings" // Package untuk memanipulasi string, seperti memisahkan teks dan mengubah huruf besar/kecil
)

// hitungSkor menghitung jumlah soal yang berhasil diselesaikan dalam waktu kurang atau sama dengan 300 menit, serta menghitung total waktu yang dibutuhkan.
func hitungSkor(waktuPengerjaan_2138 []int, jumlahSoal_2138 *int, totalWaktu_2138 *int) { // Menghitung jumlah soal dan total waktu yang valid
	*jumlahSoal_2138 = 0
	*totalWaktu_2138 = 0
	for _, waktu_2138 := range waktuPengerjaan_2138 {
		if waktu_2138 <= 300 { // Menghitung hanya soal yang diselesaikan dalam <= 300 menit
			*jumlahSoal_2138++             // Menambah jumlah soal yang berhasil diselesaikan
			*totalWaktu_2138 += waktu_2138 // Menambah waktu pengerjaan ke total waktu
		}
	}
}

// prosesInput memproses input dari pengguna dan menentukan pemenang.
func prosesInput() { // Memproses input dan menentukan pemenang berdasarkan skor
	scanner := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari pengguna

	fmt.Println(">>>>>>> PROGRAM GEMA MENCARI PEMENANG DARI DAFTAR PESERTA BY RICO ADE PRATAMA_2311102138 <<<<<<<")
	fmt.Println("Masukkan nama dan waktu pengerjaan soal (ketik 'Selesai' untuk mengakhiri input):")

	pemenang_2138, maxSoal_2138, minWaktu_2138 := "", -1, 99999 // Inisialisasi pemenang, skor maksimum, dan waktu minimum

	// Memproses input dari pengguna sampai 'Selesai' diketik
	for scanner.Scan() {
		line_2138 := scanner.Text()
		if strings.ToLower(line_2138) == "selesai" { // Jika pengguna mengetik 'Selesai', keluar dari loop
			break
		}

		data_2138 := strings.Fields(line_2138) // Memisahkan input menjadi beberapa bagian (nama dan waktu pengerjaan soal)
		if len(data_2138) != 9 {               // Memeriksa apakah input memiliki 9 bagian (1 nama dan 8 waktu)
			fmt.Println("Input tidak valid, harap masukkan nama dan 8 waktu pengerjaan soal.") // Memberikan pesan kesalahan jika input tidak valid
			continue
		}

		nama_2138 := data_2138[0]              // Mengambil nama peserta
		waktuPengerjaan_2138 := make([]int, 8) // Membuat array untuk menyimpan waktu pengerjaan soal
		for i := 0; i < 8; i++ {
			waktu_2138, error_2138 := strconv.Atoi(data_2138[i+1]) // Konversi string ke integer
			if error_2138 != nil {
				fmt.Printf("Input waktu tidak valid: %s\n", data_2138[i+1]) // Memberikan pesan kesalahan jika waktu tidak valid
				continue
			}
			waktuPengerjaan_2138[i] = waktu_2138 // Menyimpan waktu pengerjaan ke array
		}

		var jumlahSoal_2138, totalWaktu_2138 int
		hitungSkor(waktuPengerjaan_2138, &jumlahSoal_2138, &totalWaktu_2138) // Menghitung jumlah soal yang diselesaikan dan total waktu pengerjaan

		// Menentukan pemenang berdasarkan jumlah soal yang diselesaikan dan total waktu
		if jumlahSoal_2138 > maxSoal_2138 || (jumlahSoal_2138 == maxSoal_2138 && totalWaktu_2138 < minWaktu_2138) {
			pemenang_2138, maxSoal_2138, minWaktu_2138 = nama_2138, jumlahSoal_2138, totalWaktu_2138 // Update pemenang jika memenuhi kriteria
		}
	}

	// Menampilkan hasil kompetisi dan pemenang
	if pemenang_2138 == "" {
		fmt.Println("Tidak ada peserta yang valid.") // Menampilkan pesan jika tidak ada peserta yang valid
	} else {
		fmt.Println("\n<<<<<<< Pemenang Kompetisi >>>>>>>")
		fmt.Printf("Nama : %s\nJumlah soal yang diselesaikan : %d soal!\nTotal waktu : %d menit!\n", pemenang_2138, maxSoal_2138, minWaktu_2138) // Menampilkan pemenang, jumlah soal, dan total waktu
	}
}

// main adalah titik masuk utama untuk menjalankan program
func main() {
	prosesInput() // Memanggil prosedur untuk memproses input dan menentukan pemenang
}
