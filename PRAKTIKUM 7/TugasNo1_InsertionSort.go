// TUGAS 3 : PRAKTIKUM 7
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM UNTUK MEMBACA DATA INTEGER YANG KEMUDIAN DIURUTKAN MENGGUNAKAN METODE INSERTION SORT
package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk mengurutkan array_2138 secara ascending (naik) menggunakan Insertion Sort
func InsertionSort_2138(array_2138 []int) {
	n_2138 := len(array_2138) // Panjang array
	for i_2138 := 1; i_2138 < n_2138; i_2138++ {
		key_2138 := array_2138[i_2138]                     // Elemen yang sedang diurutkan
		j_2138 := i_2138 - 1                               // Indeks elemen sebelumnya
		for j_2138 >= 0 && array_2138[j_2138] > key_2138 { // Geser elemen yang lebih besar
			array_2138[j_2138+1] = array_2138[j_2138]
			j_2138--
		}
		array_2138[j_2138+1] = key_2138 // Tempatkan elemen di posisi yang benar
	}
}

// Fungsi untuk memeriksa apakah array_2138 memiliki jarak tetap
func IsJarakTetap_2138(array_2138 []int) (bool_2138 bool, jarak_2138 int) {
	if len(array_2138) < 2 { // Jika elemen kurang dari 2
		return true, 0 // Secara default dianggap jarak tetap
	}
	jarak_2138 = array_2138[1] - array_2138[0] // Hitung jarak awal
	for i_2138 := 2; i_2138 < len(array_2138); i_2138++ {
		if array_2138[i_2138]-array_2138[i_2138-1] != jarak_2138 { // Jika jarak berbeda
			return false, 0 // Tidak berjarak tetap
		}
	}
	return true, jarak_2138 // Jika jarak tetap, kembalikan true dan jarak
}

func main() {
	scanner_2138 := bufio.NewScanner(os.Stdin) // Scanner untuk membaca input

	// Meminta jumlah kasus
	fmt.Print("[2138] Masukkan jumlah kasus: ")
	scanner_2138.Scan()
	jumlahKasus_2138, _ := strconv.Atoi(scanner_2138.Text()) // Konversi jumlah kasus ke integer

	for kasus_2138 := 1; kasus_2138 <= jumlahKasus_2138; kasus_2138++ {
		// Membaca input untuk setiap kasus
		fmt.Printf(">> Masukkan data untuk kasus ke-%d (pisah dengan spasi dan akhiri dengan bilangan negatif): ", kasus_2138)
		scanner_2138.Scan()
		baris_2138 := scanner_2138.Text()         // Membaca input sebagai string
		elemen_2138 := strings.Fields(baris_2138) // Memisahkan string menjadi array
		array_2138 := make([]int, 0)              // Array untuk menyimpan bilangan non-negatif

		for _, el_2138 := range elemen_2138 { // Iterasi setiap elemen input
			num_2138, _ := strconv.Atoi(el_2138) // Konversi string ke integer
			if num_2138 < 0 {                    // Jika negatif, hentikan proses
				break
			}
			array_2138 = append(array_2138, num_2138) // Tambahkan ke array
		}

		// Urutkan array menggunakan Insertion Sort
		InsertionSort_2138(array_2138)

		// Periksa apakah array berjarak tetap
		jarakTetap_2138, jarak_2138 := IsJarakTetap_2138(array_2138)

		// Cetak hasil pengurutan
		fmt.Printf("== Hasil Pengurutan: %v\n", array_2138)
		if jarakTetap_2138 {
			fmt.Printf("=> Data berjarak %d\n\n", jarak_2138) // Jika jarak tetap, cetak jaraknya
		} else {
			fmt.Println("=> Data berjarak tidak tetap\n") // Jika tidak tetap, cetak pesan ini
		}
	}
}
