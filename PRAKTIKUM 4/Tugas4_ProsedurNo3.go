// TUGAS 4 : PRAKTIKUM 4
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM SKENA UNTUK MENCETAK SUKU DERET DARI SUKU AWAL
package main

import (
	"fmt" // Package untuk format I/O seperti print ke layar
)

// cetakDeret mencetak setiap suku dari deret bilangan hingga mencapai 1 sesuai aturan yang diberikan
func cetakDeret(n_2138 int) { // Mencetak deret bilangan mulai dari nilai n_2138
	for n_2138 != 1 {
		fmt.Printf("%d ", n_2138) // Mencetak suku saat ini diikuti spasi
		if n_2138%2 == 0 {
			n_2138 /= 2 // Jika n_2138 genap, bagi 2
		} else {
			n_2138 = 3*n_2138 + 1 // Jika n_2138 ganjil, kalikan 3 dan tambah 1
		}
	}
	fmt.Printf("1\n") // Mencetak 1 sebagai akhir deret
}

// main adalah titik masuk utama untuk menjalankan program
func main() {
	var n_2138 int
	fmt.Println("Masukkan nilai suku awal (n bilangan positif < 1000000): ") // Meminta pengguna memasukkan nilai awal
	fmt.Scan(&n_2138)

	if n_2138 > 1 && n_2138 < 1000000 { // Memastikan nilai n_2138 berada dalam batas yang valid
		cetakDeret(n_2138) // Memanggil prosedur untuk mencetak deret
	} else {
		fmt.Println("Nilai tidak valid. Harap masukkan nilai antara 1 dan 1000000.") // Menampilkan pesan kesalahan jika n_2138 tidak valid
	}
}
