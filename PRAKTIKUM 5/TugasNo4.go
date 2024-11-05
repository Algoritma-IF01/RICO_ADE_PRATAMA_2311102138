// TUGAS 4 : PRAKTIKUM 5
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM PALINDROM
package main

import (
	"bufio"   // bufio menyediakan buffered I/O, digunakan untuk membaca input dari pengguna
	"fmt"     // Mengimpor paket fmt untuk I/O
	"os"      // untuk bekerja dengan sistem operasi, seperti membaca input
	"strings" // untuk memanipulasi string, seperti perubahan huruf besar/kecil
)

const NMAX int = 127 // NMAX adalah konstanta batas dari maksimum elemen array

type tabel [NMAX]rune // tabel adalah array karakter dengan batas maksimum NMAX, dan rune berfungsi untuk menyimpan karakter unicode atau encoding karakter

// Fungsi untuk isi Array
/*
I.S. Data tersedia dalam piranti masukan
F.S. Array t berisi sejumlah n karakter yang dimasukkan user,

	Proses input selama karakter bukanlah TITIK dan n <= NMAX
*/
func isiArray(t *tabel, n *int, line string) {
	*n = 0                      // inisialisasi 0 untuk jumlah elemen array
	for _, char := range line { // untuk setiap karakter pada baris input
		if *n >= NMAX { // jika jumlah elemen array sudah mencapai batas maksimum
			break // hentikan loop
		}
		t[*n] = char // masukkan karakter ke dalam array
		*n++         // tambahkan jumlah elemen array
	}
}

// Fungsi untuk cetak Array
/*
I.S. Terdefinisi array t yang berisi sejumlah n karakter
F.S. n karakter dalam array muncul di layar
*/
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ { // untuk setiap elemen pada array
		fmt.Print(string(t[i]), " ") // cetak elemen
	}
	fmt.Println() // mencetak baris baru setelah mencetak semua elemen
}

// Fungsi untuk membalikkan isi array
/*
I.S. Terdefinisi array t yang berisi sejumlah n karakter
F.S. Urutan isi array t terbalik
*/
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ { // untuk setiap elemen pada setengah array
		t[i], t[n-1-i] = t[n-1-i], t[i] // tukar elemen
	}
}

// Fungsi utama untuk memeriksa apakah array membentuk palindrom
/*
I.S. Terdefinisi array t yang berisi sejumlah n karakter
F.S. Menghasilkan nilai true jika array t adalah palindrom, false jika tidak
*/
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ { // untuk setiap elemen pada setengah array
		if t[i] != t[n-1-i] { // jika elemen tidak sama
			return false // array tidak membentuk palindrom
		}
	}
	return true // array membentuk palindrom
}

// Fungsi utama
func main() {
	// si array tab dengan memanggil prosedur isiArray
	var tab tabel                         // Deklarasi variabel array
	var m int                             // Deklarasi variabel untuk jumlah elemen
	scanner := bufio.NewScanner(os.Stdin) // Membaca input dari pengguna
	fmt.Println("Masukkan teks (ketik '.' untuk berhenti): ")

	for scanner.Scan() { // Loop untuk membaca input pengguna
		line := scanner.Text()            // Membaca satu baris input
		if strings.ToUpper(line) == "." { // Jika input adalah titik, hentikan loop
			break // keluar dari loop
		}

		// Mengisi array dengan baris input
		isiArray(&tab, &m, line) // si array tab dengan memanggil prosedur isiArray

		// Menampilkan array asli
		fmt.Print("Teks         : ")
		cetakArray(tab, m) // Memanggil fungsi untuk mencetak array

		// Balikan isi array tab dengan memanggil balkanArray
		balikanArray(&tab, m) // Balikian isi array tab
		fmt.Print("Reverse Teks : ")
		cetakArray(tab, m) // Cetak isi array tab

		// Mengecek dan menampilkan hasil palindrom
		isPalindrom := palindrom(tab, m)           // Memeriksa apakah teks merupakan palindrom
		fmt.Println("Palindrom    ?", isPalindrom) // Menampilkan hasil palindrom
		fmt.Println()                              // mencetak baris baru untuk pemisahan
	}
}
