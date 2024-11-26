// TUGAS 1 : PRAKTIKUM 7
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM RUMAH KERABAT YANG AKAN MENYUSUN NOMOR-NOMOR RUMAH KERABATNYA SECARA TERURUT MEMBESAR MENGGUNAKAN ALGORITMA SELECTION SORT
package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk mengurutkan array_2138 secara ascending (naik) menggunakan Selection Sort
func SelectionSortAsc_2138(array_2138 []int) {
	n_2138 := len(array_2138) // Panjang array
	for i_2138 := 0; i_2138 < n_2138-1; i_2138++ {
		minIndex_2138 := i_2138 // Indeks elemen minimum
		for j_2138 := i_2138 + 1; j_2138 < n_2138; j_2138++ {
			if array_2138[j_2138] < array_2138[minIndex_2138] { // Bandingkan elemen saat ini dengan elemen minimum
				minIndex_2138 = j_2138 // Perbarui indeks elemen minimum
			}
		}
		// Tukar elemen
		array_2138[i_2138], array_2138[minIndex_2138] = array_2138[minIndex_2138], array_2138[i_2138]
	}
}

func main() {
	scanner_2138 := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari pengguna

	// Meminta jumlah daerah
	fmt.Print("[2138] Masukkan jumlah daerah: ")
	scanner_2138.Scan()
	jumlahDaerah_2138, _ := strconv.Atoi(scanner_2138.Text()) // Konversi jumlah daerah ke integer

	for i_2138 := 1; i_2138 <= jumlahDaerah_2138; i_2138++ {
		// Meminta input untuk setiap daerah
		fmt.Printf(">> Masukkan jumlah rumah (angka pertama) dan nomor rumah (pisahkan dengan spasi) untuk daerah ke-%d: ", i_2138)
		scanner_2138.Scan()
		input_2138 := scanner_2138.Text()        // Membaca input pengguna sebagai string
		parts_2138 := strings.Fields(input_2138) // Memisahkan string menjadi array string

		// Parsing jumlah rumah dan nomor rumah
		jumlahRumah_2138, _ := strconv.Atoi(parts_2138[0]) // Konversi jumlah rumah dari string ke integer
		nomorRumah_2138 := make([]int, jumlahRumah_2138)   // Buat array untuk menyimpan nomor rumah
		for j_2138 := 0; j_2138 < jumlahRumah_2138; j_2138++ {
			nomorRumah_2138[j_2138], _ = strconv.Atoi(parts_2138[j_2138+1]) // Konversi setiap nomor rumah ke integer
		}

		// Urutkan nomor rumah menggunakan Selection Sort (ascending)
		SelectionSortAsc_2138(nomorRumah_2138) // Panggil fungsi untuk mengurutkan nomor rumah

		// Tampilkan hasil untuk daerah saat ini
		fmt.Printf("== Hasil Pengurutan Nomor Rumah daerah ke-%d: ", i_2138)
		fmt.Println(strings.Trim(fmt.Sprint(nomorRumah_2138), "[]")) // Cetak nomor rumah terurut tanpa tanda kurung
	}
}
