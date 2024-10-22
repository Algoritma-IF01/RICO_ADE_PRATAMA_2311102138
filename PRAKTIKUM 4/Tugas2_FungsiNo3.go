// TUGAS 2 : PRAKTIKUM 4
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM FUNGSI UNTUK MENENTUKAN POSISI SEBUAH TITIK SEMBARANG BERADA DI SUATU LINGKARAN ATAU TIDAK
package main

import (
	"fmt"  // Package untuk format I/O seperti print ke layar
	"math" // Package untuk perhitungan matematika
)

// Fungsi untuk menghitung jarak antara dua titik (x1_2138, y1_2138) dan (x2_2138, y2_2138)
func jarak(x1_2138, y1_2138, x2_2138, y2_2138 float64) float64 {
	// Menggunakan rumus jarak Euclidean
	return math.Sqrt((x2_2138-x1_2138)*(x2_2138-x1_2138) + (y2_2138-y1_2138)*(y2_2138-y1_2138))
}

// Fungsi untuk memeriksa apakah titik (x_2138, y_2138) berada di dalam lingkaran dengan pusat (cx_2138, cy_2138) dan radius r_2138
func diDalam(cx_2138, cy_2138, r_2138, x_2138, y_2138 float64) bool {
	// Mengecek apakah jarak antara titik pusat dan titik sembarang kurang dari atau sama dengan radius
	return jarak(cx_2138, cy_2138, x_2138, y_2138) <= r_2138
}

func main() {
	var cx1_2138, cy1_2138, r1_2138 float64 // Variabel untuk pusat dan radius lingkaran 1
	var cx2_2138, cy2_2138, r2_2138 float64 // Variabel untuk pusat dan radius lingkaran 2
	var x_2138, y_2138 float64              // Variabel untuk koordinat titik sembarang

	// Input untuk lingkaran 1
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scan(&cx1_2138, &cy1_2138, &r1_2138)

	// Input untuk lingkaran 2
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scan(&cx2_2138, &cy2_2138, &r2_2138)

	// Input untuk titik sembarang
	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scan(&x_2138, &y_2138)

	// Mengecek posisi titik terhadap lingkaran 1 dan lingkaran 2
	diDalamLingkaran1_2138 := diDalam(cx1_2138, cy1_2138, r1_2138, x_2138, y_2138) // Memeriksa apakah titik berada di dalam lingkaran 1
	diDalamLingkaran2_2138 := diDalam(cx2_2138, cy2_2138, r2_2138, x_2138, y_2138) // Memeriksa apakah titik berada di dalam lingkaran 2

	// Menentukan dan menampilkan hasil
	if diDalamLingkaran1_2138 && diDalamLingkaran2_2138 {
		// Jika titik berada di dalam lingkaran 1 dan 2
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1_2138 {
		// Jika titik hanya berada di dalam lingkaran 1
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamLingkaran2_2138 {
		// Jika titik hanya berada di dalam lingkaran 2
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		// Jika titik berada di luar kedua lingkaran
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
