// TUGAS 2 : PRAKTIKUM 7
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM KERABAT DEKAT YANG AKAN MENAMPILKAN NOMOR RUMAH MULAI DARI NOMOR YANG GANJIL LEBIH DAHULU TERURUT MEMBESAR DAN KEMUDIAN MENAMPILKAN NOMOR RUMAH DENGAN NOMER GENAP TERURUT MENGECIL
package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk memisahkan bilangan ganjil dan genap dari array_2138
func PisahkanGanjilGenap_2138(array_2138 []int) (ganjil_2138 []int, genap_2138 []int) {
	for _, angka_2138 := range array_2138 {
		if angka_2138%2 == 0 {
			genap_2138 = append(genap_2138, angka_2138) // Tambahkan ke array genap_2138
		} else {
			ganjil_2138 = append(ganjil_2138, angka_2138) // Tambahkan ke array ganjil_2138
		}
	}
	return
}

// Fungsi untuk mengurutkan array_2138 secara ascending (naik) menggunakan Selection Sort
func SelectionSortAsc_2138(array_2138 []int) {
	n_2138 := len(array_2138) // Panjang array
	for i_2138 := 0; i_2138 < n_2138-1; i_2138++ {
		minIndex_2138 := i_2138 // Indeks elemen minimum
		for j_2138 := i_2138 + 1; j_2138 < n_2138; j_2138++ {
			if array_2138[j_2138] < array_2138[minIndex_2138] { // Jika elemen lebih kecil ditemukan
				minIndex_2138 = j_2138 // Perbarui indeks elemen minimum
			}
		}
		// Tukar elemen
		array_2138[i_2138], array_2138[minIndex_2138] = array_2138[minIndex_2138], array_2138[i_2138]
	}
}

// Fungsi untuk mengurutkan array_2138 secara descending (turun) menggunakan Selection Sort
func SelectionSortDesc_2138(array_2138 []int) {
	n_2138 := len(array_2138) // Panjang array
	for i_2138 := 0; i_2138 < n_2138-1; i_2138++ {
		maksIndex_2138 := i_2138 // Indeks elemen maksimum
		for j_2138 := i_2138 + 1; j_2138 < n_2138; j_2138++ {
			if array_2138[j_2138] > array_2138[maksIndex_2138] { // Jika elemen lebih besar ditemukan
				maksIndex_2138 = j_2138 // Perbarui indeks elemen maksimum
			}
		}
		// Tukar elemen
		array_2138[i_2138], array_2138[maksIndex_2138] = array_2138[maksIndex_2138], array_2138[i_2138]
	}
}

func main() {
	scanner_2138 := bufio.NewScanner(os.Stdin) // Membaca input dari pengguna baris demi baris

	fmt.Print("[2138] Masukkan jumlah daerah: ")
	scanner_2138.Scan()
	jumlahDaerah_2138, _ := strconv.Atoi(scanner_2138.Text()) // Konversi jumlah daerah ke integer

	for i_2138 := 0; i_2138 < jumlahDaerah_2138; i_2138++ {
		// Meminta input untuk setiap daerah
		fmt.Printf(">> Masukkan jumlah rumah (angka pertama) dan nomor rumah (pisahkan dengan spasi) untuk daerah ke-%d: ", i_2138+1)
		scanner_2138.Scan()
		baris_2138 := scanner_2138.Text()         // Membaca input pengguna sebagai string
		elemen_2138 := strings.Fields(baris_2138) // Memisahkan string menjadi array string

		// Parsing jumlah rumah dan nomor rumah
		jumlahRumah_2138, _ := strconv.Atoi(elemen_2138[0]) // Konversi jumlah rumah ke integer
		nomorRumah_2138 := make([]int, jumlahRumah_2138)    // Buat array untuk menyimpan nomor rumah
		for j_2138 := 0; j_2138 < jumlahRumah_2138; j_2138++ {
			nomorRumah_2138[j_2138], _ = strconv.Atoi(elemen_2138[j_2138+1]) // Konversi setiap nomor rumah ke integer
		}

		// Pisahkan angka ganjil dan genap
		ganjil_2138, genap_2138 := PisahkanGanjilGenap_2138(nomorRumah_2138)

		// Urutkan ganjil dari kecil ke besar
		SelectionSortAsc_2138(ganjil_2138)

		// Urutkan genap dari besar ke kecil
		SelectionSortDesc_2138(genap_2138)

		// Gabungkan hasil ganjil diikuti genap
		hasil_2138 := append(ganjil_2138, genap_2138...) // Gabungkan ganjil dan genap

		// Tampilkan output untuk daerah tersebut
		fmt.Printf("== Hasil Pengurutan Nomor Rumah daerah ke-%d: ", i_2138+1)
		fmt.Println(strings.Trim(fmt.Sprint(hasil_2138), "[]")) // Cetak hasil tanpa tanda kurung
	}
}
